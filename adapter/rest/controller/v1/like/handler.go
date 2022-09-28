package like

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	albumRepo "wallet-api/adapter/database/sql/album"
	likeRepo "wallet-api/adapter/database/sql/like"
	photoRepo "wallet-api/adapter/database/sql/photo"
	"wallet-api/adapter/rest/utils"
	"wallet-api/internal/dto"
	likeUseCase "wallet-api/internal/use_cases/like"
)

// @BasePath /v1

// Create - receive an objects likes , and returns a Value objects
// @Summary - give a likes
// @Description - give a likes
// @Tags - Like
// @Accept json
// @Produce json
// @Param Like body dto.LikeRequest true "Like to be created"
// @Success 201
// @Router /like [post]
// @Security ApiKeyAuth
func Create(c *gin.Context) {

	likeBody := &dto.LikeRequest{}

	err := c.BindJSON(&likeBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Error{Message: "Cannot bind like to json: " + err.Error()})
		return
	}
	idParam, found := c.Get("id")
	if !found {
		c.Status(http.StatusBadRequest)
		return
	}
	userID := idParam.(uint64)

	like := likeBody.ConvertToVO()

	if like.PhotoID != nil {

		found, _, err = likeUseCase.SaveLikePhoto(
			userID,
			like,
			photoRepo.GetPhoto,
			likeRepo.CreateLike,
			likeRepo.IncrementLikeCountPhoto,
		)
		if err != nil {
			c.JSON(http.StatusBadRequest, &dto.Error{Message: "can't likes this photo: " + err.Error()})
			return
		}
		if !found {
			c.JSON(http.StatusNotFound, &dto.Error{Message: "not found: " + err.Error()})
			return
		}
	}

	if like.AlbumID != nil {
		found, _, err = likeUseCase.SaveLikeAlbum(
			userID,
			like,
			albumRepo.GetByID,
			likeRepo.CreateLike,
			likeRepo.IncrementLikeCountAlbum,
		)
		if err != nil {
			c.JSON(http.StatusBadRequest, &dto.Error{Message: "Can't likes this album: " + err.Error()})
			return
		}
		if !found {
			c.JSON(http.StatusNotFound, &dto.Error{Message: "not found :" + err.Error()})
			return
		}
	}

	c.JSON(http.StatusCreated, nil)
}

// @BasePath /v1

// GetAllLikesByPhoto - receive a photo_id param , and returns a Value objects
// @Summary - show all likes of the photos
// @Description - show all likes of the photos
// @Tags - Like
// @Accept json
// @Produce json
// @Param photo_id path int true "photo_id"
// @Param   limit     query     int     true  "limit per page"
// @Param   page      query     int     true  "number of page"
// @Param   sort      query     string  false  "search sort"
// @Success 204
// @Router /like/photo/{photo_id} [get]
// @Security ApiKeyAuth
func GetAllLikesByPhoto(c *gin.Context) {

	pagination := utils.GeneratePaginationFromRequest(c)

	photoIDParam := c.Param("photo_id")
	photoID, err := strconv.Atoi(photoIDParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Error{Message: "ID has to be a integer: " + err.Error()})
		return
	}

	found, likes, err := likeUseCase.GetAllLikesByPhotoID(
		&pagination,
		uint64(photoID),
		likeRepo.GetLikeByPhotoID,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Error{Message: "can't get all likes: " + err.Error()})
		return
	}
	if !found {
		c.Status(http.StatusNotFound)
		return
	}

	likeDTO := &dto.Like{}
	likesDTO := likeDTO.ParseFromArrayVO(likes)

	c.JSON(http.StatusOK, &likesDTO)
}

// @BasePath /v1

// GetAllLikesByAlbum - receive an album_id param , and returns a Value objects
// @Summary - show all likes of the albums
// @Description - show all likes of the albums
// @Tags - Like
// @Accept json
// @Produce json
// @Param album_id path int true "album_id"
// @Param   limit     query     int     true  "limit per page"
// @Param   page      query     int     true  "number of page"
// @Param   sort      query     string  false  "search sort"
// @Success 204
// @Router /like/{album_id}/album [get]
// @Security ApiKeyAuth
func GetAllLikesByAlbum(c *gin.Context) {

	pagination := utils.GeneratePaginationFromRequest(c)

	albumIDParam := c.Param("album_id")
	albumID, err := strconv.Atoi(albumIDParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Error{Message: "ID has to be a integer: " + err.Error()})
		return
	}

	found, likes, err := likeUseCase.GetAllLikesByAlbum(
		&pagination,
		uint64(albumID),
		likeRepo.GetLikeByAlbumID,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Error{Message: "can't get all likes: " + err.Error()})
		return
	}

	if !found {
		c.Status(http.StatusNotFound)
		return
	}

	likeDTO := &dto.Like{}
	likesDTO := likeDTO.ParseFromArrayVO(likes)

	c.JSON(http.StatusOK, &likesDTO)
}

// @BasePath /v1

// Delete - receive a comment id , and delete this comment
// @Summary - delete a likes
// @Description - delete a likes
// @Tags - Like
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 204
// @Router /like/{id} [delete]
// @Security ApiKeyAuth
func Delete(c *gin.Context) {

	likeIDParam := c.Param("id")
	likeID, err := strconv.Atoi(likeIDParam)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	idParam, found := c.Get("id")
	if !found {
		c.Status(http.StatusBadRequest)
		return
	}
	userID := idParam.(uint64)

	err = likeUseCase.DeleteLike(
		userID,
		uint64(likeID),
		likeRepo.GetLikeByID,
		likeRepo.DeleteLike,
		likeRepo.DecrementLikeCountPhoto,
		likeRepo.DecrementLikeCountAlbum,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Error{Message: "Can't delete this like: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}
