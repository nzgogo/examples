package main

import (
	"examples/user_service/service/server"
	"log"
)

func main() {
	defer server.DB.Close()

	server.Service.Options().Transport.SetHandler(server.Service.ServerHandler)

	if err := server.Service.Init(); err != nil {
		log.Fatal(err)
	}

	initRoutes()

	if err := server.Service.Run(); err != nil {
		log.Fatal(err)
	}
}
