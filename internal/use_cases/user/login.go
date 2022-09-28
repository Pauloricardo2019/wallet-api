package user

import (
	"github.com/google/uuid"
	"time"
	tokenRepo "wallet-api/adapter/database/sql/token"
	"wallet-api/adapter/database/sql/user"
	"wallet-api/internal/error_map"
	"wallet-api/internal/model"
)

func generateUserToken(ID uint64) (*model.Token, error) {
	UserToken := &model.Token{
		Value:     uuid.New().String(),
		UserID:    ID,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(time.Hour * 24),
	}

	return UserToken, nil
}

func Login(login *model.LoginRequest, createToken tokenRepo.CreateTokenFn, getUserByEmail user.GetByEmailFn) (bool, *model.Token, error) {

	condition, userFind, err := getUserByEmail(login.Email)
	if err != nil {
		return condition, nil, error_map.WrapError(error_map.ErrValidateLogin, "User not found")
	}

	login.HashedPassword = encodePassword(login.Password)
	login.Password = ""

	if userFind.HashedPassword != login.HashedPassword {
		return false, nil, error_map.WrapError(error_map.ErrValidateLogin, "User not found")
	}

	token, err := generateUserToken(userFind.ID)
	if err != nil {
		return false, nil, error_map.WrapError(error_map.ErrValidateLogin, "Cannot generate Token")
	}

	tokenCreated, err := createToken(token)
	if err != nil {
		return false, nil, error_map.WrapError(error_map.ErrValidateLogin, "Cannot create Token")
	}

	return true, tokenCreated, nil
}
