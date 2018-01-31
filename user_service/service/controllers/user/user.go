package user

import (
	"examples/user_service/service/db"
	"examples/user_service/service/server"

	"github.com/nzgogo/micro/codec"
)

var (
	srv    = server.Service
	dbConn = server.DB.DB()
)

// type Handler func(*codec.Message, string) error

func CreateUser(msg *codec.Message, reply string) error {
	user := db.User{}

	if msg.Get("email") != "" && msg.Get("password") != "" {
		user.Email = msg.Get("email")
		user.Password = msg.Get("password")
	}

	dbConn.NewRecord(user)
	dbConn.Create(&user)

	resp := codec.NewJsonResponse(msg.ContextID, 201, user)
	srv.Respond(resp, reply)

	return nil
}

func GetUser(msg *codec.Message, reply string) error {
	if msg.Get("email") != "" {
		user := db.User{}
		dbConn.Where(&db.User{
			Email: msg.Get("email"),
		}).First(&user)

		if user.ID == "" {
			srv.Respond(
				codec.NewJsonResponse(msg.ContextID, 404, "User not found!"),
				reply,
			)
		} else {
			srv.Respond(
				codec.NewJsonResponse(msg.ContextID, 200, user),
				reply,
			)
		}
	} else {
		srv.Respond(
			codec.NewJsonResponse(msg.ContextID, 422, "Email not presented in request!"),
			reply,
		)
	}

	return nil
}

func UpdateUser(msg *codec.Message, reply string) error {
	return nil
}

func DeleteUser(msg *codec.Message, reply string) error {
	return nil
}
