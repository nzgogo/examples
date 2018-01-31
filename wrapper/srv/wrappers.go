package main

import (
	"log"
	"github.com/nzgogo/micro/codec"
	"github.com/nzgogo/micro/router"

)

func logMsgWrapper(handler router.Handler) router.Handler {
	return func(msg *codec.Message, reply string) error {
		//log.Printf("logMsgwrapper -> message received: %v\n" , *msg)
		log.Printf("[1] logMsgwrapper before")
		err := handler(msg, reply)
		log.Printf("[1] logMsgwrapper after")
		return err
	}
}

func logMsgWrapper2(handler router.Handler) router.Handler {
	return func(msg *codec.Message, reply string) error {
		//log.Printf("logMsgwrapper -> message received: %v\n" , *msg)
		log.Printf("[2] logMsgwrapper before")
		err := handler(msg, reply)
		log.Printf("[2] logMsgwrapper after")
		return err
	}
}

func logMsgWrapper3(handler router.Handler) router.Handler {
	return func(msg *codec.Message, reply string) error {
		//log.Printf("logMsgwrapper -> message received: %v\n" , *msg)
		log.Printf("[3] logMsgwrapper before")
		err := handler(msg, reply)
		log.Printf("[3] logMsgwrapper after")
		return err
	}
}
