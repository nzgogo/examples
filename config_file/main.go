package main

import (
	"log"

	"github.com/nzgogo/micro"
)

const (
	SrvName    = "gogo-test-config"
	SrvVersion = "v1"
)

func main() {
	srv := gogo.NewService(SrvName, SrvVersion)

	log.Println("-----Config Start-----")

	for k, v := range srv.Config() {
		log.Printf("[%s]%s\n", k, v)
	}

	log.Println("-----Config End-----")

	if err := srv.Init(); err != nil {
		log.Fatal(err)
	}

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
