package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"runtime/debug"

	jsoniter "github.com/json-iterator/go"
	"github.com/nzgogo/govalidator"
	"github.com/nzgogo/micro"
	"github.com/nzgogo/micro/codec"
	"github.com/nzgogo/micro/constant"
	"github.com/nzgogo/micro/context"
	recpro "github.com/nzgogo/micro/recover"
	"github.com/nzgogo/micro/router"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type MyHandler struct {
	srv gogo.Service
}

func init() {
	customRuleObjectId()
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

	//look up registered service in kv store
	var node *router.Node
	node, err = config.Router.HttpMatch(r.URL.Path, r.Method)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// read request body and then reset body before processing validation
	var bodyBytes []byte
	if r.Body != nil && strings.Contains(r.Header.Get("Content-Type"), "application/json") {
		bodyBytes, _ = ioutil.ReadAll(r.Body)
		r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	}

	//input validation
	if node.ValidationRules != nil && len(node.ValidationRules) > 0 {
		opts := govalidator.Options{
			Request:  r,
			Rules:    node.ValidationRules,
			Messages: node.ValidationMessages,
		}
		if strings.Contains(r.Header.Get("Content-Type"), "application/json") {
			validationData := make(map[string]interface{}, 0)
			opts.Data = &validationData
			v := govalidator.New(opts)
			e := v.ValidateJSON()
			if len(e) > 0 {
				errorMessage := map[string]interface{}{"validation error": e}
				w.Header().Set("Content-type", "application/json")
				w.WriteHeader(http.StatusUnprocessableEntity)
				json.NewEncoder(w).Encode(errorMessage)
				return
			}
		} else {
			v := govalidator.New(opts)
			e := v.Validate()
			if len(e) > 0 {
				errorMessage := map[string]interface{}{"validation error": e}
				w.Header().Set("Content-type", "application/json")
				w.WriteHeader(http.StatusUnprocessableEntity)
				json.NewEncoder(w).Encode(errorMessage)
				return
			}
		}
	}

	// create context and map the HTTP request to internal transport request message struct.
	ctxId := config.Context.Add(&context.Conversation{
		Response: w,
	})
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
	log.Println("Parsed service: " + srvName + "-" + srvVersion)

	//service discovery
	subj, err := config.Selector.Select(srvName, srvVersion)
	if err != nil {
		http.Error(w, "Cannot process request", http.StatusInternalServerError)
		//panic("None service available. " + err.Error())
		return
	}
	log.Println("Found service: " + subj)

	//transport file
	//reqBody := make(map[string]interface{})
	//codec.Unmarshal(request.Body, &reqBody)
	reqBody := request.Body
	if reqBody["file"] != nil {
		log.Println("File found!")
		file, ok := reqBody["file"].(*codec.File)
		if !ok {
			http.Error(w, "Cannot process request", http.StatusInternalServerError)
			panic("Failed to parse file from request body. ")
		}
		if fileSub, err := config.Selector.Select(constant.FILE_SERVICE_NAME, constant.FILE_SERVICE_VERSION); err != nil {
			http.Error(w, "Cannot process request", http.StatusInternalServerError)
			panic("None service available. " + err.Error())
		} else {
			h.srv.Options().Transport.SendFile(request, fileSub, file.Data)
			log.Println("Send file to transport.")
		}
		file.Data = nil
		request.Set("file", file)
	}

	// transport request
	bytes, _ := codec.Marshal(request)
	log.Println("Send to service: " + subj)
	respErr := config.Transport.Publish(subj, bytes)

	if respErr != nil {
		http.Error(w, "Transport error", http.StatusInternalServerError)
		panic("Transport error, Published failed. " + respErr.Error())
	}

	config.Context.Wait(ctxId)
}

func main() {
	service := gogo.NewService(
		"gogo-core-api",
		"v1",
	)
	service.Options().Transport.SetHandler(service.ApiHandler)

	var respWrapChain = []gogo.HttpResponseWrapper{
		logHttpRespWrapper,
		logHttpRespWrapper2,
		logHttpRespWrapper3,
	}

	if err := service.Init(gogo.WrapRepsWriter(respWrapChain...)); err != nil {

		log.Fatal(err)
	}

	go func() {
		if err := service.Run(); err != nil {
			log.Fatal(err)
		}
	}()

	handler := MyHandler{service}
	server := http.Server{
		Addr: "0.0.0.0:8080",
		//Handler: &handler,
	}
	// Http wrapper
	var httpChain = []gogo.HttpHandlerWrapper{
		handler.logWrapper,
	}
	http.HandleFunc("/", gogo.HttpWrapperChain(handler.ServeHTTP, httpChain...))
	go func() {
		if err := server.ListenAndServe(); err != nil {
			// cannot panic, because this probably is an intentional close
			log.Printf("Httpserver: ListenAndServe() error: %s", err)
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
