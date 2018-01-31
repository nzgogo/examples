package main

import (
	"net/http"
	"log"
	"fmt"
	"github.com/nzgogo/micro"
)

func logwrapper(wrapper gogo.HttpHandlerFunc) gogo.HttpHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("before")
		wrapper(w,r)
		log.Println("after")
	}
}



// AuthMiddleware - takes in a http.HandlerFunc, and returns a http.HandlerFunc
var AuthMiddleware = func(f http.HandlerFunc) http.HandlerFunc {
	// one time scope setup area for middleware
	return func(w http.ResponseWriter, r *http.Request) {
		// ... pre handler functionality
		fmt.Println("start auth")
		f(w, r)
		fmt.Println("end auth")
		// ... post handler functionality
	}
}

// PrivateMiddleware - takes in a http.HandlerFunc, and returns a http.HandlerFunc
var PrivateMiddleware = func(f http.HandlerFunc) http.HandlerFunc {
	// one time scope setup area for middleware
	return func(w http.ResponseWriter, r *http.Request) {
		// ... pre handler functionality
		fmt.Println("start private")
		f(w, r)
		fmt.Println("end private")
		// ... post handler functionality
	}
}
