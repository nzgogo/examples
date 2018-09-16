package controllers

import (
	"examples/prometheus_custom_metrics/db"
	"examples/prometheus_custom_metrics/globals"
	"github.com/nzgogo/micro/codec"
	"github.com/nzgogo/micro/router"
	"net/http"
)

// type Handler func(*codec.Message, string) error

func CreateUser(msg *codec.Message, reply string) *router.Error {
	user := db.User{}
	if err := user.Insert(); err != nil {
		return &router.Error{StatusCode: http.StatusInternalServerError, Message: err.Error()}
	}

	resp := codec.NewJsonResponse(msg.ContextID, http.StatusCreated)
	globals.ContextCnt.Desc()
	if err := globals.Service.Respond(resp, reply); err != nil {
		return &router.Error{StatusCode: http.StatusInternalServerError, Message: err.Error()}
	}

	return nil
}
