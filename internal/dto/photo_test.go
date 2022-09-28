package dto_test

import (
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"wallet-api/internal/dto"
	"wallet-api/internal/model"
)

func TestPhoto_ParseFromVO(t *testing.T) {
	fakerGen := faker.New()
	var urlimage = fakerGen.Internet().URL()

	photo := &model.Photo{
		ID:           fakerGen.UInt64(),
		AlbumID:      fakerGen.UInt64(),
		UrlImage:     &urlimage,
		CreatedAt:    time.Time{},
		DeletedAt:    time.Now(),
		LikeCount:    fakerGen.UInt64(),
		CommentCount: fakerGen.UInt64(),
	}

	dtoPhoto := &dto.Photo{}

	dtoPhoto.ParseFromVO(photo)

	assert.Equal(t, photo.ID, dtoPhoto.ID)
	assert.Equal(t, photo.AlbumID, dtoPhoto.AlbumID)
	assert.Equal(t, photo.UrlImage, dtoPhoto.UrlImage)
	assert.Equal(t, photo.CreatedAt, dtoPhoto.CreatedAt)
	assert.Equal(t, photo.DeletedAt, dtoPhoto.DeletedAt)
	assert.Equal(t, photo.LikeCount, dtoPhoto.LikeCount)
	assert.Equal(t, photo.CommentCount, dtoPhoto.CommentCount)
}

func TestPhoto_ConvertToVO(t *testing.T) {
	fakerGen := faker.New()
	var urlimage = fakerGen.Internet().URL()

	dtoPhoto := dto.Photo{
		ID:           fakerGen.UInt64(),
		AlbumID:      fakerGen.UInt64(),
		UrlImage:     &urlimage,
		CreatedAt:    time.Time{},
		DeletedAt:    time.Now(),
		LikeCount:    fakerGen.UInt64(),
		CommentCount: fakerGen.UInt64(),
	}

	photo := dtoPhoto.ConvertToVO()

	assert.Equal(t, photo.ID, dtoPhoto.ID)
	assert.Equal(t, photo.AlbumID, dtoPhoto.AlbumID)
	assert.Equal(t, photo.UrlImage, dtoPhoto.UrlImage)
	assert.Equal(t, photo.CreatedAt, dtoPhoto.CreatedAt)
	assert.Equal(t, photo.DeletedAt, dtoPhoto.DeletedAt)
	assert.Equal(t, photo.LikeCount, dtoPhoto.LikeCount)
	assert.Equal(t, photo.CommentCount, dtoPhoto.CommentCount)
}

func TestPhoto_ConvertToArrayDTO(t *testing.T) {
	fakerGen := faker.New()
	var urlimage = fakerGen.Internet().URL()

	dtoPhoto1 := dto.Photo{
		ID:           fakerGen.UInt64(),
		AlbumID:      fakerGen.UInt64(),
		UrlImage:     &urlimage,
		CreatedAt:    time.Time{},
		DeletedAt:    time.Now(),
		LikeCount:    fakerGen.UInt64(),
		CommentCount: fakerGen.UInt64(),
	}

	dtoPhoto2 := dto.Photo{
		ID:           fakerGen.UInt64(),
		AlbumID:      fakerGen.UInt64(),
		UrlImage:     &urlimage,
		CreatedAt:    time.Time{},
		DeletedAt:    time.Now(),
		LikeCount:    fakerGen.UInt64(),
		CommentCount: fakerGen.UInt64(),
	}
	dtoPhoto3 := dto.Photo{
		ID:           fakerGen.UInt64(),
		AlbumID:      fakerGen.UInt64(),
		UrlImage:     &urlimage,
		CreatedAt:    time.Time{},
		DeletedAt:    time.Now(),
		LikeCount:    fakerGen.UInt64(),
		CommentCount: fakerGen.UInt64(),
	}

	dtoPhotosBeforeFunc := []dto.Photo{
		dtoPhoto1,
		dtoPhoto2,
		dtoPhoto3,
	}

	dtoPhotos := []dto.Photo{
		dtoPhoto1,
		dtoPhoto2,
		dtoPhoto3,
	}

	var photos []model.Photo

	for i := 0; i < len(photos); i++ {
		dtoPhotos[i].ConvertToArrayDTO(photos)

		assert.Equal(t, dtoPhotos[i].ID, dtoPhotosBeforeFunc[i].ID)
		assert.Equal(t, dtoPhotos[i].AlbumID, dtoPhotosBeforeFunc[i].AlbumID)
		assert.Equal(t, dtoPhotos[i].UrlImage, dtoPhotosBeforeFunc[i].UrlImage)
		assert.Equal(t, dtoPhotos[i].CreatedAt, dtoPhotosBeforeFunc[i].CreatedAt)
		assert.Equal(t, dtoPhotos[i].DeletedAt, dtoPhotosBeforeFunc[i].DeletedAt)
		assert.Equal(t, dtoPhotos[i].LikeCount, dtoPhotosBeforeFunc[i].LikeCount)
		assert.Equal(t, dtoPhotos[i].CommentCount, dtoPhotosBeforeFunc[i].CommentCount)

	}

}
