package model

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type UserProfile struct {
	ID        uint64    `gorm:"primaryKey"`
	UrlImage  *string   `valid:"-"`
	FirstName string    `valid:"required" gorm:"varchar(100)"`
	LastName  string    `valid:"required" gorm:"varchar(100)"`
	CreatedAt time.Time `valid:"-" `
	UpdatedAt time.Time `valid:"-" `
}

func (u *UserProfile) Validate() error {
	return validator.New().Struct(u)
}

func (UserProfile) Tablename() string {
	return "user_profile"
}
