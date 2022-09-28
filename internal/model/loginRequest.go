package model

type LoginRequest struct {
	Email          string `valid:"notnull" gorm:"varchar(255)"`
	Password       string `valid:"notnull" gorm:"varchar(255)"`
	HashedPassword string `valid:"notnull" gorm:"-"`
}
