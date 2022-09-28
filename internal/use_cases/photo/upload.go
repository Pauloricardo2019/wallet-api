package photo

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	s3 "wallet-api/adapter/aws/s3"
	albumRepo "wallet-api/adapter/database/sql/album"
	photoRepo "wallet-api/adapter/database/sql/photo"
	"wallet-api/internal/error_map"
	"wallet-api/internal/model"

	"github.com/google/uuid"
)

func SaveUpload(userID uint64, albumID uint64, getAlbum albumRepo.GetByIDFn, file *multipart.FileHeader, savePhoto photoRepo.CreatePhotoFn) error {

	found, album, err := getAlbum(albumID)
	if err != nil {
		return error_map.WrapError(error_map.ErrValidationAlbum, "error to upload photo")
	}

	if !found {
		return error_map.WrapError(error_map.ErrValidationAlbum, "error to upload photo")
	}

	if album == nil {
		return error_map.WrapError(error_map.ErrValidationAlbum, "error to upload photo")
	}

	if album.UserID != userID {
		return error_map.WrapError(error_map.ErrValidationAlbum, "error to upload photo")
	}

	albumIDInt := int(albumID)
	userIDInt := int(userID)
	albumIDStrg := strconv.Itoa(albumIDInt)
	userIDStrg := strconv.Itoa(userIDInt)

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

	urlPhoto, err := s3.UploadAlbumFile(userIDStrg, albumIDStrg, fileName, fileData)
	if err != nil {
		return err
	}
	photo := &model.Photo{}
	photo.UrlImage = urlPhoto
	photo.AlbumID = albumID

	err = savePhoto(photo)
	if err != nil {
		return err
	}

	return nil
}
