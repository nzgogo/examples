package main

import (
	"examples/user_service/service/db"
	"examples/user_service/service/server"
	"log"

	"github.com/nzgogo/micro"
)

const (
	SrvName    = "gogox-core-hello"
	SrvVersion = "v1"
)

func main() {
	srv := gogo.NewService(SrvName, SrvVersion)

	server.Service = srv
	server.DB = db.NewDB()

	defer server.DB.Close()

	server.Service.Options().Transport.SetHandler(srv.ServerHandler)

	if err := server.Service.Init(); err != nil {
		log.Fatal(err)
	}

	if err := server.Service.Run(); err != nil {
		log.Fatal(err)
	}
}
