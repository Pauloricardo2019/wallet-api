package user

import (
	userRepo "wallet-api/adapter/database/sql/user"
	"wallet-api/internal/model"
)

func GetByID(userID uint64, getByID userRepo.GetByIDFn) (bool, *model.User, error) {

	condition, user, err := getByID(userID)
	if err != nil {
		return condition, nil, err
	}

	return condition, user, nil
}
