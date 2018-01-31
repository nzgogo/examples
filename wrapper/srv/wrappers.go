package main

import (
	"log"
	"github.com/nzgogo/micro/codec"
	"github.com/nzgogo/micro/router"

)

func logMsgwrapper(handler router.Handler) router.Handler {
	return func(msg *codec.Message, reply string) error {
		log.Printf("logMsgwrapper -> message received: %v\n" , *msg)
		err := handler(msg, reply)
		return err
	}
}
