package photo

import (
	photoRepo "wallet-api/adapter/database/sql/photo"
	"wallet-api/internal/model"
)

func GetByID(photoID uint64, getPhotoByID photoRepo.GetPhotoFn) (bool, *model.Photo, error) {
	found, photo, err := getPhotoByID(photoID)

	if err != nil {
		return false, nil, err
	}

	if !found {
		return false, nil, err
	}

	if photo == nil {
		return false, nil, err
	}

	return true, photo, nil
}
