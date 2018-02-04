package user

import (
	"examples/user_service/service/db"
	"examples/user_service/service/server"

	"github.com/nzgogo/micro/codec"
	"github.com/nzgogo/micro/router"
)

var (
	srv    = server.Service
	dbConn = server.DB.DB()
)

// type Handler func(*codec.Message, string) error

func CreateUser(msg *codec.Message, reply string) *router.Error {
	user := db.User{}

	if msg.Get("email") != "" && msg.Get("password") != "" {
		user.Email = msg.Get("email")
		user.Password = msg.Get("password")
	}

	dbConn.NewRecord(user)
	dbConn.Create(&user)

	resp := codec.NewJsonResponse(msg.ContextID, 201, user)
	server.Service.Respond(resp, reply)

	return nil
}

func GetUser(msg *codec.Message, reply string) *router.Error {
	if msg.Query.Get("email") != "" {
		user := db.User{}
		dbConn.Where(&db.User{
			Email: msg.Query.Get("email"),
		}).First(&user)

		if user.ID == "" {
			return &router.Error{StatusCode: 404, Message: "User not found!"}
		} else {
			server.Service.Respond(
				codec.NewJsonResponse(msg.ContextID, 200, user),
				reply,
			)
		}
	} else {
		return &router.Error{StatusCode: 422, Message: "Email not presented in request!"}
	}

	return nil
}

func UpdateUser(msg *codec.Message, reply string) *router.Error {
	if msg.Get("id") != "" {
		user := db.User{}
		dbConn.Where(&db.User{ID: msg.Get("id")}).First(&user)

		values := msg.GetAll()
		for k, v := range values {
			if v == "" {
				continue
			}
			switch k {
			case "email":
				user.Email = v
			case "password":
				user.Password = v
			}
		}
		dbConn.Save(&user)

		server.Service.Respond(
			codec.NewJsonResponse(msg.ContextID, 200, user),
			reply,
		)
	} else {
		return &router.Error{StatusCode: 422, Message: "User ID not presented in request!"}
	}
	return nil
}

func DeleteUser(msg *codec.Message, reply string) *router.Error {
	if msg.Get("id") != "" {
		user := db.User{ID: msg.Get("id")}
		dbConn.Delete(&user)

		server.Service.Respond(
			codec.NewJsonResponse(msg.ContextID, 200, ""),
			reply,
		)
	} else {
		return &router.Error{StatusCode: 422, Message: "User ID not presented in request!"}
	}
	return nil
}
