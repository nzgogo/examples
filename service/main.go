package main

import (
	"github.com/nats-io/go-nats"
	"github.com/nzgogo/micro"
)

func main() {
	srv := gogo.NewService(
		"gogox.core.hello",
		"v1",
	)

	srv.Options().Transport.SetHandler(func(msg *nats.Msg) {

	})

	srv.Init()

	srv.Run()
}
