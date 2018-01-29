package main

import (
	"log"

	"github.com/nzgogo/micro"
	"github.com/nzgogo/micro/db"
)

type Server struct {
	srv gogo.Service
	db  db.DB
}

func main() {
	srv := gogo.NewService(
		"gogox-core-hello",
		"v1",
	)

	server := Server{
		srv: srv,
		db: db.NewDB(
			"gogo",
			"gogox123",
			"test",
			db.Address("gogo-api-test.c69ll9boyxmw.ap-southeast-2.rds.amazonaws.com"),
		),
	}

	defer server.db.Close()

	server.srv.Options().Transport.SetHandler(srv.ServerHandler)

	if err := server.srv.Init(); err != nil {
		log.Fatal(err)
	}

	if err := server.srv.Run(); err != nil {
		log.Fatal(err)
	}
}
