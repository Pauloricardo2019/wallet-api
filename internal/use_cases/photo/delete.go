package photo

import (
	albumRepo "wallet-api/adapter/database/sql/album"
	photoRepo "wallet-api/adapter/database/sql/photo"
	"wallet-api/internal/model"
)

func DeletePhotoByID(userID uint64, photoID uint64, getPhoto photoRepo.GetPhotoFn, deletedPhoto photoRepo.DeletePhotoFn, saveDeleted photoRepo.SaveDeletedPhotoFn, getAlbum albumRepo.GetByIDFn) error {

	PhotoDeleted := &model.DeletedPhoto{}

	found, photo, err := getPhoto(photoID)
	if err != nil {
		return err
	}
	if !found {
		return err
	}

	found, album, err := getAlbum(photo.AlbumID)
	if err != nil {
		return err
	}
	if !found {
		return err
	}

	if userID != album.UserID {
		return err
	}

	PhotoDeleted.UserID = userID
	PhotoDeleted.AlbumID = album.ID
	PhotoDeleted.PhotoID = photoID

	err = saveDeleted(PhotoDeleted)
	if err != nil {
		return err
	}

	err = deletedPhoto(photoID)
	if err != nil {
		return err
	}

	return nil

}
