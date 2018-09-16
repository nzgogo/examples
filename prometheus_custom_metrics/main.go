package main

import (
	"examples/prometheus_custom_metrics/globals"
	"github.com/nzgogo/micro"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

func main() {
	globals.Service.Options().Transport.SetHandler(globals.Service.ServerHandler)
	// config wrapper ContextMetricWrapper
	if err := globals.Service.Init(gogo.WrapHandler(ContextMetricWrapper)); err != nil {
		log.Fatal(err)
	}
	// register endpoints to kv store
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
