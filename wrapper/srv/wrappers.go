package main

import (
	"log"
	"net/http"
	"github.com/nzgogo/micro/codec"
	"github.com/nzgogo/micro/router"
	"github.com/nzgogo/micro"
)

func logMsgwrapper(handler router.Handler) router.Handler {
	return func(msg *codec.Message, reply string) error {
		log.Printf("logMsgwrapper -> message received: %v\n" , *msg)
		err := handler(msg, reply)
		return err
	}
}

func logHttpRespwrapper(writeResponse gogo.HttpResponseWriter) gogo.HttpResponseWriter  {
	return func(rw http.ResponseWriter, response *codec.Message) {
		log.Printf("logHttpRespwrapper -> message to send to http ResponseWriter: %v\n" , *response)
		writeResponse(rw, response)
	}
}