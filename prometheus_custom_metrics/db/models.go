package db

import "github.com/nzgogo/mgo/bson"

type User struct {
	ID   bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Name string        `bson:"name,omitempty" json:"name,omitempty"`
}

func (user *User) Insert() error {
	return nil
}
