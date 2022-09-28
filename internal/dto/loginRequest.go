package dto

import "wallet-api/internal/model"

type LoginRequest struct {
	Email    string `valid:"notnull" json:"email"`
	Password string `valid:"notnull" json:"password"`
} // @name LoginRequest

func (dto *LoginRequest) ConvertToVO() *model.LoginRequest {
	result := &model.LoginRequest{}

	result.Email = dto.Email
	result.Password = dto.Password

	return result
}
