package db

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	ID        string `gorm:"type:varchar(36);primary_key"`
	Email     string `gorm:"type:varchar(255);unique"`
	Password  string `gorm:"size:255"`
	Roles     []Role `gorm:"many2many:user_roles;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Role struct {
	gorm.Model
	Name        string `gorm:"size:255"`
	Description string `gorm:"size:1024"`
	Users       []User `gorm:"many2many:user_roles;"`
}
