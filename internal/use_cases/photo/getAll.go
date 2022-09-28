package photo

import (
	"strconv"
	"wallet-api/adapter/aws/s3"
	albumRepo "wallet-api/adapter/database/sql/album"
	photoRepo "wallet-api/adapter/database/sql/photo"
	"wallet-api/internal/model"
)

func GetAllPhotos(pagination *model.Pagination, albumID uint64, getAllPhotos photoRepo.GetAllPhotosFn, getAlbum albumRepo.GetByIDFn) (bool, []model.Photo, error) {

	found, album, err := getAlbum(albumID)
	if err != nil {
		return false, nil, err
	}
	if !found {
		return false, nil, err
	}

	found, photos, err := getAllPhotos(album.ID, pagination)
	if err != nil {
		return false, nil, err
	}
	if !found {
		return false, nil, err
	}

	for _, v := range photos {
		if v.AlbumID != albumID {
			return false, nil, err
		}
	}

	userIDInt := int(album.UserID)
	albumIDInt := int(albumID)
	userIDStrg := strconv.Itoa(userIDInt)
	albumIDStrg := strconv.Itoa(albumIDInt)

	err = s3.ListFiles(photos, userIDStrg, albumIDStrg)
	if err != nil {
		return false, nil, err
	}

	return true, photos, nil
}
