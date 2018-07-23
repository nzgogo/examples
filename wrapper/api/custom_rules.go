package main

import (
	"fmt"

	"github.com/nzgogo/govalidator"
	"github.com/nzgogo/mgo/bson"
)

func customRuleObjectId() {
	govalidator.AddCustomRule("objectid", func(field string, rule string, message string, value interface{}) error {
		v := string(fmt.Sprintf("%v", value))
		if !bson.IsObjectIdHex(v) {
			return fmt.Errorf("The %s field must be an ObjectId", field)
		}
		return nil
	})
}