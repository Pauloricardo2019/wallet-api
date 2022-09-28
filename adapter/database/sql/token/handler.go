package token

import (
	"errors"
	"gorm.io/gorm"
	"wallet-api/adapter/database/sql"
	"wallet-api/internal/model"
)

type CreateTokenFn func(token *model.Token) (*model.Token, error)

func CreateToken(token *model.Token) (*model.Token, error) {
	db, err := sql.GetGormDB()
	if err != nil {
		return nil, err
	}

	if err = db.Create(token).Error; err != nil {
		return nil, err
	}

	return token, nil
}

type GetTokenByValueFn func(token string) (bool, *model.Token, error)

func GetTokenByValue(token string) (bool, *model.Token, error) {
	db, err := sql.GetGormDB()
	if err != nil {
		return false, nil, err
	}
	result := &model.Token{}
	if err = db.First(&result, "value = ? ", token).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, result, nil
		}
		return false, nil, err
	}

	return true, result, err

}
