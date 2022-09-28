package user

import (
	userRepo "wallet-api/adapter/database/sql/user"
	"wallet-api/internal/error_map"
)

func Delete(tokenUserID, userId uint64, deleteFn userRepo.DeleteFn) error {

	if tokenUserID != userId {
		return error_map.WrapError(error_map.ErrValidationUser, "You're not authorized to manage this user")
	}

	err := deleteFn(userId)
	if err != nil {
		return err
	}

	return nil
}
