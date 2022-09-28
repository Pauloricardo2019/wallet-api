package photo

import (
	"errors"
	"gorm.io/gorm"
	"time"
	"wallet-api/adapter/database/sql"
	"wallet-api/internal/model"
)

type CreatePhotoFn func(photo *model.Photo) error

func CreatePhoto(photo *model.Photo) error {
	db, err := sql.GetGormDB()
	if err != nil {
		return err
	}

	photo.CreatedAt = time.Now()
	if err = db.Create(photo).Error; err != nil {
		return err
	}

	return nil
}

type GetPhotoFn func(photoID uint64) (bool, *model.Photo, error)

func GetPhoto(photoID uint64) (bool, *model.Photo, error) {
	db, err := sql.GetGormDB()
	if err != nil {
		return false, nil, err
	}
	photo := &model.Photo{}
	if err = db.First(&photo, "id = ?", photoID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, photo, nil
		}
		return false, nil, err
	}

	return true, photo, nil

}

type GetAllPhotosFn func(albumID uint64, pagination *model.Pagination) (bool, []model.Photo, error)

func GetAllPhotos(albumID uint64, pagination *model.Pagination) (bool, []model.Photo, error) {
	db, err := sql.GetGormDB()
	if err != nil {
		return false, nil, err
	}

	result := []model.Photo{}

	offset := (pagination.Page - 1) * pagination.Limit

	queryBuider := db.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)

	if err := queryBuider.Where(&model.Photo{AlbumID: albumID}).Find(&result).Error; err != nil {
		return false, nil, err
	}

	return len(result) > 0, result, nil
}

type DeletePhotoFn func(photoID uint64) error

func DeletePhoto(photoID uint64) error {
	db, err := sql.GetGormDB()
	if err != nil {
		return err
	}

	if err = db.Delete(&model.Photo{ID: photoID}).Error; err != nil {
		return err
	}

	return nil
}

type UpdatePhotoFn func(photo *model.Photo) error

func UpdatePhoto(photo *model.Photo) error {
	db, err := sql.GetGormDB()
	if err != nil {
		return err
	}

	if err = db.Save(&photo).Error; err != nil {
		return err
	}

	return nil

}

type SaveDeletedPhotoFn func(deletedPhoto *model.DeletedPhoto) error

func SaveDeletedPhoto(deletedPhoto *model.DeletedPhoto) error {
	db, err := sql.GetGormDB()
	if err != nil {
		return err
	}
	deletedPhoto.DeletedAt = time.Now()
	if err = db.Create(deletedPhoto).Error; err != nil {
		return err
	}

	return nil
}
