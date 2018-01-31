package db

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID        string `gorm:"size:36;primary_key"`
	Email     string `gorm:"size:255;unique"`
	Password  string `gorm:"size:255"`
	Roles     []Role `gorm:"many2many:user_roles;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (u *User) BeforeCreate(scope *gorm.Scope) error {
	for {
		if id, err := uuid.NewV4(); err == nil {
			scope.SetColumn("ID", id.String())
			break
		}
	}
	return nil
}

type Role struct {
	gorm.Model
	Name        string `gorm:"size:255"`
	Description string `gorm:"size:1024"`
	Users       []User `gorm:"many2many:user_roles;"`
}
