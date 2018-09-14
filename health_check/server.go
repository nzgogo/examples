package main

import (
	"log"
	"net/http"
	"github.com/nzgogo/micro"
	"github.com/nzgogo/micro/codec"
	"github.com/nzgogo/micro/router"
)

type server struct {
	srv gogo.Service
}

var (
	responsecode = "For the brave souls who get this far: You are the chosen ones, the valiant knights of programming who toil away, without rest, fixing our most awful code. To you, true saviors, kings of men, I say this: never gonna give you up, never gonna let you down, never gonna run around and desert you. Never gonna make you cry, never gonna say goodbye. Never gonna tell a lie and hurt you."
)

func (s *server) Hello(req *codec.Message, reply string) *router.Error  {
	//fmt.Println("Message received: " + string(req.Body))
	h := http.Header{}
	h.Add("Content-Type", "text/plain")
	response := codec.NewResponse( req.ContextID, 200, []byte(responsecode), h)
	err := s.srv.Respond(response, reply)
	if err !=nil {
		return &router.Error{
			500,
			err.Error(),
		}
	}

	return nil
}

func main() {
	server := server{}
	service := gogo.NewService(
		"gogo-core-greeter",
		"v1",
	)

	server.srv = service

	if err := server.srv.Init(); err != nil {
		log.Fatal(err)
	}

	server.srv.Options().Transport.SetHandler(service.ServerHandler)

	r := server.srv.Options().Router

	r.Add(&router.Node{
		Method:  "GET",
		Path:    "/hello",
		ID:      "/hello",
		Handler: server.Hello,
	})

	// Run server
	if err := server.srv.Run(); err != nil {
		log.Fatal(err)
	}
}
