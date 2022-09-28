//go:build integration_test
// +build integration_test

package photo_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math"
	"testing"
	albumRepo "wallet-api/adapter/database/sql/album"
	photoRepo "wallet-api/adapter/database/sql/photo"
	userRepo "wallet-api/adapter/database/sql/user"
	"wallet-api/internal/model"
)

func TestCreatePhoto(t *testing.T) {

	user := &model.User{
		FirstName: "Test12",
	}

	user, err := userRepo.Create(user)
	require.NoError(t, err)

	album := &model.Album{
		UserID: user.ID,
		Name:   "AlbumTest",
	}
	album, err = albumRepo.Create(album)
	require.NoError(t, err)

	photo := &model.Photo{
		AlbumID: album.ID,
	}
	err = photoRepo.CreatePhoto(photo)
	require.NoError(t, err)

}

func TestGetPhoto(t *testing.T) {

	user := &model.User{
		FirstName: "Test12",
	}

	user, err := userRepo.Create(user)
	require.NoError(t, err)

	album := &model.Album{
		UserID: user.ID,
		Name:   "AlbumTest",
	}
	album, err = albumRepo.Create(album)
	require.NoError(t, err)

	photo := &model.Photo{
		AlbumID: album.ID,
	}
	err = photoRepo.CreatePhoto(photo)
	require.NoError(t, err)

	found, _, err := photoRepo.GetPhoto(photo.ID)
	require.NoError(t, err)

	assert.True(t, found)

}

func TestGetPhoto_NotFound(t *testing.T) {

	found, _, err := photoRepo.GetPhoto(math.MaxInt32)
	require.NoError(t, err)
	assert.False(t, found)

}

func TestUpdatePhoto(t *testing.T) {

	user := &model.User{
		FirstName: "Test12",
	}

	user, err := userRepo.Create(user)
	require.NoError(t, err)

	album := &model.Album{
		UserID: user.ID,
		Name:   "AlbumTest",
	}
	album, err = albumRepo.Create(album)
	require.NoError(t, err)

	photo := &model.Photo{
		AlbumID: album.ID,
	}
	err = photoRepo.CreatePhoto(photo)
	require.NoError(t, err)

	err = photoRepo.UpdatePhoto(photo)
	require.NoError(t, err)

}

func TestGetAllPhotos(t *testing.T) {
	user := &model.User{
		FirstName: "Test12",
	}

	user, err := userRepo.Create(user)
	require.NoError(t, err)

	album := &model.Album{
		UserID: user.ID,
		Name:   "AlbumTest",
	}
	album, err = albumRepo.Create(album)
	require.NoError(t, err)

	photo := &model.Photo{
		AlbumID: album.ID,
	}
	err = photoRepo.CreatePhoto(photo)
	require.NoError(t, err)

	pagination := &model.Pagination{
		Limit: 10,
		Page:  1,
		Sort:  "created_at asc",
	}

	found, _, err := photoRepo.GetAllPhotos(photo.AlbumID, pagination)
	require.NoError(t, err)

	assert.True(t, found)

}

func TestGetAllPhotos_NotFound(t *testing.T) {

	pagination := &model.Pagination{
		Limit: 10,
		Page:  1,
		Sort:  "created_at asc",
	}

	found, _, err := photoRepo.GetAllPhotos(math.MaxInt32, pagination)
	require.NoError(t, err)
	assert.False(t, found)

}

func TestDeletePhoto(t *testing.T) {

	user := &model.User{
		FirstName: "Test12",
	}

	user, err := userRepo.Create(user)
	require.NoError(t, err)

	album := &model.Album{
		UserID: user.ID,
		Name:   "AlbumTest",
	}
	album, err = albumRepo.Create(album)
	require.NoError(t, err)

	photo := &model.Photo{
		AlbumID: album.ID,
	}
	err = photoRepo.CreatePhoto(photo)
	require.NoError(t, err)

	err = photoRepo.DeletePhoto(photo.ID)
	require.NoError(t, err)

}

func TestSaveDeletedPhoto(t *testing.T) {
	user := &model.User{
		FirstName: "Test12",
	}

	user, err := userRepo.Create(user)
	require.NoError(t, err)

	album := &model.Album{
		UserID: user.ID,
		Name:   "AlbumTest",
	}
	album, err = albumRepo.Create(album)
	require.NoError(t, err)

	photo := &model.Photo{
		AlbumID: album.ID,
	}

	err = photoRepo.CreatePhoto(photo)
	require.NoError(t, err)

	deletedPhoto := &model.DeletedPhoto{
		PhotoID: photo.ID,
		AlbumID: album.ID,
		UserID:  user.ID,
	}

	err = photoRepo.SaveDeletedPhoto(deletedPhoto)
	require.NoError(t, err)

}
