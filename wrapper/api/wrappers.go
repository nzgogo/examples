package main

import (
	"net/http"
	"log"
	"github.com/nzgogo/micro"
	"github.com/nzgogo/micro/codec"
)

func logWrapper(wrapper gogo.HttpHandlerFunc) gogo.HttpHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("[1] HttpHandlerWrapper before")
		wrapper(w,r)
		log.Println("[1] HttpHandlerWrapper after")
	}
}

func logWrapper2(wrapper gogo.HttpHandlerFunc) gogo.HttpHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("[2] HttpHandlerWrapper")
		wrapper(w,r)
		log.Println("[2] HttpHandlerWrapper after")
	}
}

func logWrapper3(wrapper gogo.HttpHandlerFunc) gogo.HttpHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("[3] HttpHandlerWrapper before")
		wrapper(w,r)
		log.Println("[3] HttpHandlerWrapper after")
	}
}


func logHttpRespWrapper(writeResponse gogo.HttpResponseWriter) gogo.HttpResponseWriter  {
	return func(rw http.ResponseWriter, response *codec.Message) {
		//log.Printf("[1] logHttpRespwrapper -> message to send to http ResponseWriter: %s\n" , response.Body)
		log.Printf("[1] logHttpRespwrapper before")
		writeResponse(rw, response)
		log.Printf("[1] logHttpRespwrapper after")
	}
}

func logHttpRespWrapper2(writeResponse gogo.HttpResponseWriter) gogo.HttpResponseWriter  {
	return func(rw http.ResponseWriter, response *codec.Message) {
		//log.Printf("[2] logHttpRespwrapper -> message to send to http ResponseWriter: %s\n" , response.Body)
		log.Printf("[2] logHttpRespwrapper before")
		writeResponse(rw, response)
		log.Printf("[2] logHttpRespwrapper after")
	}
}

func logHttpRespWrapper3(writeResponse gogo.HttpResponseWriter) gogo.HttpResponseWriter  {
	return func(rw http.ResponseWriter, response *codec.Message) {
		//log.Printf("[3] logHttpRespwrapper -> message to send to http ResponseWriter: %s\n" , response.Body)
		log.Printf("[3] logHttpRespwrapper before")
		writeResponse(rw, response)
		log.Printf("[3] logHttpRespwrapper after")
	}
}

// AuthMiddleware - takes in a http.HandlerFunc, and returns a http.HandlerFunc
//var AuthMiddleware = func(f http.HandlerFunc) http.HandlerFunc {
//	// one time scope setup area for middleware
//	return func(w http.ResponseWriter, r *http.Request) {
//		// ... pre handler functionality
//		fmt.Println("start auth")
//		f(w, r)
//		fmt.Println("end auth")
//		// ... post handler functionality
//	}
//}
