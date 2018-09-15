package main

import (
	"examples/prometheus_custom_metrics/controllers"
	"examples/prometheus_custom_metrics/globals"
	"log"

	"github.com/nzgogo/micro/router"
)

var routes []*router.Node

func init() {
	routes = append(routes, &router.Node{
		Method:  "POST",
		Path:    "/user",
		ID:      "create_single_user",
		Handler: controllers.CreateUser,
	})
}

func initRoutes() {
	log.Println("-----Service Routes List Start-----")
	for _, route := range routes {
		globals.Service.Options().Router.Add(route)
		log.Printf("%v\n", route)
	}
	log.Println("-----Service Routes List End-----")
}
