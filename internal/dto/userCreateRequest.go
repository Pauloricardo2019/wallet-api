package dto

import (
	"wallet-api/internal/model"
)

type UserCreateRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password,omitempty"`
} // @name UserCreateRequest

func (dto *UserCreateRequest) GenerateUserVO() *model.User {
	result := &model.User{}

	result.FirstName = dto.FirstName
	result.LastName = dto.LastName
	result.Email = dto.Email
	result.Username = dto.Username
	result.Password = dto.Password

	return result
}
