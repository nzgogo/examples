package main

import (
	"bytes"
	"github.com/nzgogo/micro"
	"github.com/nzgogo/micro/codec"
	"github.com/nzgogo/micro/constant"
	"github.com/nzgogo/micro/context"
	recpro "github.com/nzgogo/micro/recover"
	"github.com/nzgogo/micro/router"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
)

const (
	SRVNAME          = "gogo-core-api"
	SRVVERSION       = "v1"
	SRVADDR          = "0.0.0.0:8080"
)

type MyHandler struct {
	srv gogo.Service
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if rMsg := recover(); rMsg != nil {
			recpro.PostProc(h.srv.Config()[constant.SLACKCHANNELADDR], h.srv.Options().Transport.Options().Subject, "ServeHTTP", rMsg, string(debug.Stack()))
			http.Error(w, "Cannot process request", http.StatusInternalServerError)
		}
	}()

	var err error
	config := h.srv.Options()

	//look up the registered services stored in consul kv store
	var node *router.Node
	node, err = config.Router.HttpMatch(r.URL.Path, r.Method)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// read request body and then reset body before processing validation
	var bodyBytes []byte
	if r.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(r.Body)
		r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	}

	// create context
	ctxId := config.Context.Add(&context.Conversation{
		Response: w,
	})
	// map the HTTP request to internal message struct.
	jsonData := make(map[string]interface{})
	if len(bodyBytes) > 0 {
		codec.Unmarshal(bodyBytes, &jsonData)
	}
	request := &codec.Message{}
	request.Node = node.ID
	request, err = request.ParseHTTPRequest(r, config.Transport.Options().Subject, ctxId, jsonData)
	if err != nil {
		http.Error(w, "Cannot process request", http.StatusNotFound)
		return
	}

	//parse service name and version
	srvName := gogo.URLToServiceName(request.Host, request.Path)
	srvVersion := gogo.URLToServiceVersion(request.Path)

	//service discovery
	subj, err := config.Selector.Select(srvName, srvVersion)
	if err != nil {
		http.Error(w, "Cannot process request", http.StatusInternalServerError)
		return
	}

	// transport request
	reqBytes, _ := codec.Marshal(request)
	log.Println("Send to service: " + subj)

	respErr := config.Transport.Publish(subj, reqBytes)
	if respErr != nil {
		http.Error(w, "Transport error", http.StatusInternalServerError)
		panic("Transport error, Published failed. " + respErr.Error())
	}

	config.Context.Wait(ctxId)
}

func main() {
	service := gogo.NewService(
		SRVNAME,
		SRVVERSION,
	)
	service.Options().Transport.SetHandler(service.ApiHandler)

	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := service.Run(); err != nil {
			log.Fatal(err)
		}
	}()

	handler := MyHandler{service}
	server := http.Server{
		Addr: SRVADDR,
	}
	http.HandleFunc("/", handler.ServeHTTP)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			// cannot panic, because this probably is an intentional close
			log.Fatal("HttpServer: ListenAndServe() error: " + err.Error())
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	select {
	// wait on kill signal
	case <-ch:
		service.Stop()
		if err := server.Shutdown(nil); err != nil {
			panic(err) // failure/timeout shutting down the server gracefully
		}
	}
}

