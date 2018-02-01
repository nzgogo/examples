package main

import (
	"log"
	"net/http"

	"github.com/nzgogo/micro"
)

var httpRequestWrappers = []gogo.HttpHandlerWrapper{
	logWrapper,
}

func logWrapper(wrapper gogo.HttpHandlerFunc) gogo.HttpHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("-----HTTP Request Received-----")
		log.Println("Proto: " + r.Proto)
		log.Println("Method: " + r.Method)
		log.Println("URI: " + r.RequestURI)
		wrapper(w, r)
	}
}
