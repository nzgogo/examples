package main

import (
	"examples/prometheus_custom_metrics/globals"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

func main() {
	globals.Service.Options().Transport.SetHandler(globals.Service.ServerHandler)
	if err := globals.Service.Init(); err != nil {
		log.Fatal(err)
	}
	initRoutes()

	// export metrics to prometheus server through port 2112
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		if err := http.ListenAndServe(":2112", nil); err != nil {
			log.Fatal("HttpServer: ListenAndServe() error: " + err.Error())
		}
	}()

	// start service
	if err := globals.Service.Run(); err != nil {
		log.Fatal(err)
	}
}
