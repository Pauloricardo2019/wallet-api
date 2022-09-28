package model

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type User struct {
	ID             uint64       `gorm:"primaryKey"`
	Email          string       `valid:"required" gorm:"varchar(255)"`
	Birth          time.Time    `valid:"-"`
	Password       string       `valid:"required" gorm:"-"`
	HashedPassword string       `valid:"required" gorm:"varchar(64)"`
	UserProfileID  *uint64      `valid:"-"`
	UserProfile    *UserProfile `valid:"-" gorm:"foreign_key:UserProfileID"`
	CreatedAt      time.Time    `valid:"-" `
	UpdatedAt      time.Time    `valid:"-" `
}

func (u *User) Validate() error {
	return validator.New().Struct(u)
}

func (User) Tablename() string {
	return "user"
}
