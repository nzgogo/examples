package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/nzgogo/micro"
	"github.com/nzgogo/micro/api"
	"github.com/nzgogo/micro/codec"
	"github.com/nzgogo/micro/context"
)

type APIHTTPHandler struct {
	srv gogo.Service
}

func (s *APIHTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	config := s.srv.Options()

	ctxId := config.Context.Add(&context.Conversation{
		Response: w,
	})

	// map the HTTP request to internal transport request message struct.
	request, err := gogoapi.HTTPReqToIntrlSReq(r, config.Transport.Options().Subject, ctxId)
	if err != nil {
		http.Error(w, "Cannot process request", http.StatusInternalServerError)
		return
	}

	//look up registered service in kv store
	err = config.Router.HttpMatch(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	srvName := gogo.URLToIntnlTrans(request.Host, request.Path)
	fmt.Println("Dispatch to server: " + srvName)

	//service discovery
	subj, err := config.Selector.Select(srvName, "v1")
	if err != nil {
		fmt.Printf("Selector failed. error: %v", err)
		http.Error(w, "Cannot process request", http.StatusInternalServerError)
		return
	}
	fmt.Println("Found service: " + subj)

	//transport
	bytes, _ := codec.Marshal(request)
	fmt.Println("send to service: " + subj)
	respErr := config.Transport.Publish(subj, bytes)
	if respErr != nil {
		fmt.Printf("failed to send message . error: %v", err)
		http.Error(w, "No response", http.StatusInternalServerError)
		return
	}

	config.Context.Wait(ctxId)
}

func main() {
	srv := gogo.NewService(
		"gogo-core-api",
		"v1",
	)

	srv.Options().Transport.SetHandler(srv.ApiHandler)

	if err := srv.Init(); err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := srv.Run(); err != nil {
			log.Fatal(err)
		}
	}()

	h := APIHTTPHandler{srv}
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/", gogo.HttpWrapperChain(h.ServeHTTP, httpRequestWrappers...))

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Printf("Httpserver: ListenAndServe() error: %s", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)

	select {
	case <-ch:
		if err := server.Shutdown(nil); err != nil {
			panic(err)
		}
	}
}
