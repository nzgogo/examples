package main

import (
	"fmt"
	"github.com/nzgogo/micro"
	"github.com/nzgogo/micro/codec"
	"github.com/nzgogo/micro/router"
	"log"
	"net/http"
)

var (
	Service gogo.Service
)

func Hello(req *codec.Message, reply string) *router.Error {
	fmt.Println("received a http request:")
	response := codec.NewResponse(req.ContextID, http.StatusOK)
	response.Set("data", "Hello World")
	if err := Service.Respond(response, reply); err != nil {
		return &router.Error{
			500,
			err.Error(),
		}
	}
	return nil
}

func main() {
	// create a new micro service
	Service = gogo.NewService(
		"gogo-core-greeter",
		"v1",
	)
	if err := Service.Init(); err != nil {
		log.Fatal(err)
	}
	Service.Options().Transport.SetHandler(Service.ServerHandler)

	// add hello endpoint
	r := Service.Options().Router
	r.Add(&router.Node{
		Method:  "GET",
		Path:    "/hello",
		ID:      "hello",
		Handler: Hello,
	})

	// Run service
	if err := Service.Run(); err != nil {
		log.Fatal(err)
	}
}
