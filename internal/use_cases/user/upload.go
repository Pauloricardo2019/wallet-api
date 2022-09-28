package user

import (
	"fmt"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"wallet-api/adapter/aws/s3"
	userRepo "wallet-api/adapter/database/sql/user"
	"wallet-api/internal/error_map"
)

func Upload(
	tokenUserID uint64,
	id uint64,
	getByUserID userRepo.GetByIDFn,
	updateUser userRepo.UpdateFn,
	file *multipart.FileHeader,
) error {

	if id != tokenUserID {
		return error_map.WrapError(error_map.ErrValidationUser, "error to upload photo")
	}

	found, user, err := getByUserID(id)
	if err != nil {
		return err
	}
	if !found {
		return error_map.WrapError(error_map.ErrValidationUser, "error user not found")
	}

	//Open the file
	src, err := file.Open()
	if err != nil {
		return err
	}

	//Returns an array of the bytes
	srcBytes, err := io.ReadAll(src)
	if err != nil {
		return err
	}

	fileExtension := filepath.Ext(file.Filename)

	dst := uuid.New().String()

	//FileName
	fileName := fmt.Sprintf("%s%s", dst, fileExtension)

	//Copy the byte array into the file
	err = os.WriteFile(fileName, srcBytes, 0777)
	if err != nil {
		return err
	}

	//Reads the file and returns an array of the bytes
	fileData, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	urlPhoto, err := s3.UploadUserFile(fileName, fileData)
	if err != nil {
		return err
	}

	user.UrlImage = urlPhoto

	err = updateUser(user)
	if err != nil {
		return err
	}

	return nil
}
