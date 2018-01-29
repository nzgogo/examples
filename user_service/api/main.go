package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/nzgogo/micro"
)

type HttpHandler struct {
	srv gogo.Service
}

func (h *HttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}

func main() {
	srv := gogo.NewService(
		"gogox-core-api",
		"v1",
	)

	srv.Options().Transport.SetHandler(srv.ApiHandler)

	if err := srv.Init(); err != nil {
		log.Fatal(err)
	}

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}

	h := HttpHandler{srv}
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &h,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Printf("Httpserver: ListenAndServe() error: %s", err)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)

	select {
	case <-ch:
		if err := server.Shutdown(nil); err != nil {
			panic(err)
		}
	}
}
