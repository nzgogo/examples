package server

import (
	"examples/user_service/service/db"

	"github.com/nzgogo/micro"
)

const (
	SrvName    = "gogo-core-hello"
	SrvVersion = "v1"
)

var (
	Service = gogo.NewService(SrvName, SrvVersion)
	DB      = db.NewDB()
)
