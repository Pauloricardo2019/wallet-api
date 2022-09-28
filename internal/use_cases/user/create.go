package user

import (
	"crypto/sha256"
	"fmt"
	"regexp"
	"wallet-api/adapter/database/sql/user"
	"wallet-api/internal/error_map"
	"wallet-api/internal/model"

	emailverifier "github.com/AfterShip/email-verifier"
)

func encodePassword(password string) string {
	sum := sha256.Sum256([]byte(password))
	return fmt.Sprintf("%x", sum)
}

func validatePassword(password string) error {

	if password == "" {
		return error_map.WrapError(error_map.ErrValidationUser, "password cannot to be empty")
	}

	tests := []string{".{7,}", "([a-z]{1,})", "([A-Z]{1,})", "([0-9]{1,})", "([!@#$&*]{1,})"}
	for index, test := range tests {
		t, err := regexp.MatchString(test, password)
		if err != nil {
			return error_map.WrapError(error_map.ErrValidationUser, err.Error())
		}
		if !t {
			if index == 0 {
				return error_map.WrapError(error_map.ErrValidationUser, "this password must be 7 characters long")
			}
			if index == 1 {
				return error_map.WrapError(error_map.ErrValidationUser, "this password must be at least 1 letter")
			}
			if index == 2 {
				return error_map.WrapError(error_map.ErrValidationUser, "this password must be at least 1 uppercase letter")
			}
			if index == 3 {
				return error_map.WrapError(error_map.ErrValidationUser, "this password must be at least 1 number")
			}
			if index == 4 {
				return error_map.WrapError(error_map.ErrValidationUser, "this password must be at least 1 special character")
			}
		}
	}

	return nil
}

func validateUser(user *model.User) error {

	if user.FirstName == "" {
		return error_map.WrapError(error_map.ErrValidationUser, "field FirstName cannot to be empty")
	}

	if len(user.FirstName) < 4 {
		return error_map.WrapError(error_map.ErrValidationUser, "FirstName cannot be less than 4 characters")
	}

	if len(user.FirstName) > 100 {
		return error_map.WrapError(error_map.ErrValidationUser, "FirstName cannot be longer than 100 characters")
	}

	if user.LastName == "" {
		return error_map.WrapError(error_map.ErrValidationUser, "field LastName cannot to be empty")
	}

	if len(user.LastName) < 4 {
		return error_map.WrapError(error_map.ErrValidationUser, "LastName cannot be less than 4 characters")
	}

	if len(user.LastName) > 100 {
		return error_map.WrapError(error_map.ErrValidationUser, "LastName cannot be longer than 100 characters")
	}

	if user.Email == "" {
		return error_map.WrapError(error_map.ErrValidationUser, "field email cannot to be empty")
	}

	verifier := emailverifier.NewVerifier()
	result, err := verifier.Verify(user.Email)
	if err != nil {
		return error_map.WrapError(error_map.ErrValidationUser, " verify email failed")
	}

	if !result.Syntax.Valid {
		fmt.Println("email address syntax is invalid")
		return error_map.WrapError(error_map.ErrValidationUser, " invalid email")
	}

	if user.Username == "" {
		return error_map.WrapError(error_map.ErrValidationUser, "field auth cannot to be empty")
	}

	if len(user.Username) < 4 {
		return error_map.WrapError(error_map.ErrValidationUser, "auth cannot be less than 4 characters")
	}

	if len(user.Username) > 50 {
		return error_map.WrapError(error_map.ErrValidationUser, "auth cannot be longer than 50 characters")
	}

	err = validatePassword(user.Password)
	if err != nil {
		return err
	}

	return nil
}

func Create(user *model.User, createUser user.CreateFn, getUserByEmail user.GetByEmailFn, getUserByLogin user.GetByLoginFn) (*model.User, error) {

	if err := validateUser(user); err != nil {
		return nil, err
	}

	_, existingEmail, err := getUserByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	if existingEmail.ID != 0 {
		return nil, error_map.WrapError(error_map.ErrValidationUser, "email already exists")
	}

	_, existingLogin, err := getUserByLogin(user.Username)
	if err != nil {
		return nil, error_map.WrapError(error_map.ErrValidationUser, "cannot get auth by id")
	}

	if existingLogin.ID != 0 {
		return nil, error_map.WrapError(error_map.ErrValidationUser, "auth already exists")
	}

	user.HashedPassword = encodePassword(user.Password)
	user.Password = ""

	createdUser, err := createUser(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

// comentario
