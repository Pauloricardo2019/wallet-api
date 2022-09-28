package dto

import (
	"time"
	"wallet-api/internal/model"
)

type User struct {
	ID        uint64    `json:"id" valid:"-"`
	UrlImage  *string   `valid:"-"`
	FirstName string    `json:"first_name" valid:"notnull" `
	LastName  string    `json:"last_name" valid:"notnull" `
	Email     string    `json:"email" valid:"notnull" `
	DDD       string    `json:"ddd" valid:"-"`
	Phone     string    `json:"phone" valid:"-"`
	Username  string    `json:"username" valid:"notnull" `
	Birth     time.Time `json:"birth" valid:"-"`
	Biography string    `json:"biography" valid:"-" `
	Password  string    `json:"password,omitempty" valid:"notnull"`
	CreatedAt time.Time `json:"created_at" valid:"-"`
	UpdatedAt time.Time `json:"updated_at" valid:"-"`
} // @name User

func (dto *User) ParseFromVO(user *model.User) {
	dto.ID = user.ID
	dto.UrlImage = user.UrlImage
	dto.FirstName = user.FirstName
	dto.LastName = user.LastName
	dto.Email = user.Email
	dto.DDD = user.DDD
	dto.Phone = user.Phone
	dto.Username = user.Username
	dto.Birth = user.Birth
	dto.Biography = user.Biography
	dto.Password = user.Password
	dto.CreatedAt = user.CreatedAt
	dto.UpdatedAt = user.UpdatedAt
}

func (dto *User) ConvertToVO() *model.User {
	result := &model.User{}

	result.ID = dto.ID
	result.UrlImage = dto.UrlImage
	result.FirstName = dto.FirstName
	result.LastName = dto.LastName
	result.Email = dto.Email
	result.DDD = dto.DDD
	result.Phone = dto.Phone
	result.Username = dto.Username
	result.Birth = dto.Birth
	result.Biography = dto.Biography
	result.Password = dto.Password
	result.CreatedAt = dto.CreatedAt
	result.UpdatedAt = dto.UpdatedAt

	return result
}

func (dto *User) ParseFromArrayVO(users []model.User) []User {
	var usersDTO []User

	for _, user := range users {
		usersDTO = append(usersDTO, User{
			ID:        user.ID,
			UrlImage:  user.UrlImage,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			DDD:       user.DDD,
			Phone:     user.Phone,
			Username:  user.Username,
			Birth:     user.Birth,
			Biography: user.Biography,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}
	return usersDTO
}
