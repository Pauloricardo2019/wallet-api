package token

import (
	tokenRepo "wallet-api/adapter/database/sql/token"
)

func ValidateToken(token string, getTokenByValue tokenRepo.GetTokenByValueFn) (bool, uint64) {

	found, tokenFind, err := getTokenByValue(token)
	if err != nil {
		return false, 0
	}
	if !found {
		return false, 0
	}

	if tokenFind.UserID <= 0 {
		return false, 0
	}

	return true, tokenFind.UserID
}
