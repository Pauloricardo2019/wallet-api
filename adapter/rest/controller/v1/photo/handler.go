package photo

import (
	"net/http"
	"strconv"
	albumRepo "wallet-api/adapter/database/sql/album"
	photoRepo "wallet-api/adapter/database/sql/photo"
	"wallet-api/adapter/rest/utils"
	"wallet-api/internal/dto"
	photoUseCase "wallet-api/internal/use_cases/photo"

	"github.com/gin-gonic/gin"
)

// Upload - does upload a photo, and save the path_file on database
// @Summary - Upload photo
// @Description - Upload a photo
// @Tags - Photo
// @Accept multipart/form-data
// @Produce json
// @Param album_id path int true "album_id"
// @Param file formData file true  "upload your file"
// @Success 204
// @Router /photo/{album_id} [post]
// @Security ApiKeyAuth
func Upload(c *gin.Context) {

	albumIDParam := c.Param("album_id")
	albumID, err := strconv.Atoi(albumIDParam)
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

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, &dto.Error{Message: "Cannot get a photo: " + err.Error()})
		return
	}

	err = photoUseCase.SaveUpload(userID, uint64(albumID), albumRepo.GetByID, file, photoRepo.CreatePhoto)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Error{Message: "Cannot upload photo: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, http.StatusNoContent)
}

// GetPhoto - takes a photo by id and returns it
// @Summary - get photo
// @Description - get photo by id
// @Tags - Photo
// @Accept json
// @Produce json
// @Param photo_id path int true "photo_id"
// @Success 200 {object} dto.Photo
// @Router /photo/album/{photo_id} [get]
// @Security ApiKeyAuth
func GetPhoto(c *gin.Context) {
	photoIDParam := c.Param("photo_id")
	photoID, err := strconv.Atoi(photoIDParam)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	_, found := c.Get("id")
	if !found {
		c.Status(http.StatusBadRequest)
		return
	}

	found, photo, err := photoUseCase.GetByID(uint64(photoID), photoRepo.GetPhoto)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Error{Message: "Cannot get photo: " + err.Error()})
		return
	}
	if !found {
		c.Status(http.StatusNotFound)
		return
	}

	photoDTO := &dto.Photo{}
	photoDTO.ParseFromVO(photo)

	c.JSON(http.StatusOK, &photoDTO)

}

// GetAllPhotosByAlbum - get an album by id and return an array of photos
// @Summary - get all photos
// @Description - get all photos by album
// @Tags - Photo
// @Accept json
// @Produce json
// @Param album_id path int true "album_id"
// @Param   limit     query     int     true  "limit per page"
// @Param   page      query     int     true  "number of page"
// @Param   sort      query     string  false  "search sort"
// @Success 200 {array} dto.Photo
// @Router /photo/{album_id} [get]
// @Security ApiKeyAuth
func GetAllPhotosByAlbum(c *gin.Context) {

	pagination := utils.GeneratePaginationFromRequest(c)

	albumIDParam := c.Param("album_id")
	albumID, err := strconv.Atoi(albumIDParam)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	_, found := c.Get("id")
	if !found {
		c.Status(http.StatusBadRequest)
		return
	}

	found, photos, err := photoUseCase.GetAllPhotos(&pagination, uint64(albumID), photoRepo.GetAllPhotos, albumRepo.GetByID)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Error{Message: "Cannot get all photos: " + err.Error()})
		return
	}

	if !found {
		c.Status(http.StatusNotFound)
		return
	}

	photoDTO := &dto.Photo{}
	photosDTO := photoDTO.ConvertToArrayDTO(photos)

	c.JSON(http.StatusOK, &photosDTO)
}

// DeletePhoto - delete a photo by id
// @Summary - delete a photo
// @Description - delete a photo
// @Tags - Photo
// @Accept json
// @Produce json
// @Param photo_id path int true "photo_id"
// @Success 204
// @Router /photo/{photo_id} [delete]
// @Security ApiKeyAuth
func DeletePhoto(c *gin.Context) {
	photoIDParam := c.Param("photo_id")
	photoID, err := strconv.Atoi(photoIDParam)
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

	err = photoUseCase.DeletePhotoByID(userID, uint64(photoID), photoRepo.GetPhoto, photoRepo.DeletePhoto, photoRepo.SaveDeletedPhoto, albumRepo.GetByID)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Error{Message: "Cannot delete photo: " + err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
