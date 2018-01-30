package db

import "github.com/nzgogo/micro/db"

const (
	Username = "gogo"
	Password = "gogox123"
	Database = "test"
)

func NewDB() db.DB {
	conn := db.NewDB(
		Username,
		Password,
		Database,
		db.Address("gogo-api-test.c69ll9boyxmw.ap-southeast-2.rds.amazonaws.com"),
	)

	conn.Connect()

	conn.DB().AutoMigrate(&User{})
	conn.DB().AutoMigrate(&Role{})

	return conn
}
