package model

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type Category struct {
	ID        uint64    `gorm:"primaryKey"`
	Name      string    `validate:"required"`
	CreatedAt time.Time `valid:"-" `
	UpdatedAt time.Time `valid:"-" `
}

func (c *Category) Validator() error {
	return validator.New().Struct(c)
}

func (Category) Tablename() string {
	return "category"
}
