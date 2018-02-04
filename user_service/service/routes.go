package main

import (
	"examples/user_service/service/controllers/user"
	"examples/user_service/service/server"
	"log"

	"github.com/nzgogo/micro/router"
)

var routes []*router.Node

func init() {
	routes = append(routes, &router.Node{
		Method:  "GET",
		Path:    "/user",
		ID:      "get_single_user",
		Handler: user.GetUser,
	})

	routes = append(routes, &router.Node{
		Method:  "POST",
		Path:    "/user",
		ID:      "create_single_user",
		Handler: user.CreateUser,
	})

	routes = append(routes, &router.Node{
		Method:  "PUT",
		Path:    "/user",
		ID:      "update_single_user",
		Handler: user.UpdateUser,
	})

	routes = append(routes, &router.Node{
		Method:  "DELETE",
		Path:    "/user",
		ID:      "delete_single_user",
		Handler: user.DeleteUser,
	})
}

func initRoutes() {
	log.Println("-----Service Routes List Start-----")
	for _, route := range routes {
		server.Service.Options().Router.Add(route)
		log.Printf("%v\n", route)
	}
	log.Println("-----Service Routes List End-----")
}
