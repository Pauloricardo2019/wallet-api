package error_map

import (
	"errors"
	"fmt"
)

var (
	ErrValidationUser    = errors.New("invalid user")
	ErrValidationAlbum   = errors.New("invalid album")
	ErrValidateLogin     = errors.New("invalid auth")
	ErrValidatePhoto     = errors.New("invalid photo")
	ErrValidationLike    = errors.New("invalid like")
	ErrValidationComment = errors.New("invalid comment")
)

func WrapError(err error, msg string) error {
	return errors.New(fmt.Sprintf("%s. %s", err.Error(), msg))
}
