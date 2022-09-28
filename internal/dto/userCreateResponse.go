package dto

import (
	"wallet-api/internal/model"
)

type UserCreateResponse struct {
	ID        uint64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Username  string `json:"username"`
} // @name UserCreateResponse

func (dto *UserCreateResponse) ParseFromUserVO(user *model.User) {
	dto.ID = user.ID
	dto.FirstName = user.FirstName
	dto.LastName = user.LastName
	dto.Email = user.Email
	dto.Username = user.Username
}
