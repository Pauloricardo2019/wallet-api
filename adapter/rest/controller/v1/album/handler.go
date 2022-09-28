package album

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	albumRepo "wallet-api/adapter/database/sql/album"
	photoRepo "wallet-api/adapter/database/sql/photo"
	"wallet-api/adapter/rest/utils"
	"wallet-api/internal/dto"
	albumUseCase "wallet-api/internal/use_cases/album"
)

// @BasePath /v1

// GetByID - receives id as a parameter, and returns an album objects
// @Summary - Get one album
// @Description - Get an album by id
// @Tags - Album
// @Produce json
// @Accept json
// @Param id path int true "id"
// @Success 200 {object} dto.Album
// @Router /album/{id} [get]
// @Security ApiKeyAuth
func GetByID(c *gin.Context) {

	albumID := c.Param("id")
	intAlbumID, err := strconv.Atoi(albumID)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Error{Message: "ID has to a be integer: " + err.Error()})
		return
	}

	_, found := c.Get("id")
	if !found {
		c.Status(http.StatusBadRequest)
		return
	}

	found, album, err := albumUseCase.GetById(
		uint64(intAlbumID),
		albumRepo.GetByID,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Error{Message: "Cannot find album: " + err.Error()})
		return
	}

	if !found {
		c.Status(http.StatusNotFound)
		return
	}

	albumDTO := &dto.Album{}
	albumDTO.ParseFromVO(album)

	c.JSON(http.StatusOK, albumDTO)
}

// @BasePath /v1

// GetAlbumByUserID - receives user_id as a parameter, and returns an array the albums objects of this user
// @Summary - Get albums
// @Description - Get all albums of the user
// @Tags - Album
// @Accept json
// @Produce json
// @Param user_id path int true "user_id"
// @Param   limit     query     int     true  "limit per page"
// @Param   page      query     int     true  "number of page"
// @Param   sort      query     string  false  "search sort"
// @Success 200 {array} dto.Album
// @Router /album/user/{user_id} [get]
// @Security ApiKeyAuth
func GetAlbumByUserID(c *gin.Context) {

	pagination := utils.GeneratePaginationFromRequest(c)

	userId := c.Param("user_id")
	userIDParam, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Error{Message: "ID has to a be integer: " + err.Error()})
		return
	}

	_, found := c.Get("id")
	if !found {
		c.Status(http.StatusBadRequest)
		return
	}

	found, albums, err := albumUseCase.GetAlbumsByUser(&pagination, uint64(userIDParam), albumRepo.GetAlbumByUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Error{Message: "Cannot get album by userid: " + err.Error()})
		return
	}

	if !found {
		c.Status(http.StatusNotFound)
		return
	}
	albumDTO := &dto.Album{}
	albumsDTO := albumDTO.ParseFromArrayVO(albums)
	c.JSON(http.StatusOK, &albumsDTO)
}

// @BasePath /v1

// Create - create an album, and returns an album objects
// @Summary - Create album
// @Description - Create an album
// @Tags - Album
// @Accept json
// @Produce json
// @Param Album body dto.Album true "Album to be created"
// @Success 200 {object} dto.Album
// @Router /album [post]
// @Security ApiKeyAuth
func Create(c *gin.Context) {

	paramID, found := c.Get("id")
	if !found {
		c.Status(http.StatusBadRequest)
		return
	}

	userID := paramID.(uint64)

	albumDTO := &dto.Album{}

	err := c.BindJSON(albumDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Error{Message: "Cannot bind album to json: " + err.Error()})
		return
	}

	albumVO := albumDTO.ConvertToVO()

	album, err := albumUseCase.Create(userID, albumVO, albumRepo.Create)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Error{Message: "Cannot create an album: " + err.Error()})
		return
	}

	albumDTO.ParseFromVO(album)

	c.JSON(http.StatusOK, albumDTO)
}

// @BasePath /v1

// UploadAlbumCover - upload an album cover
// @Summary - upload an album cover
// @Description - upload an album cover
// @Tags - Album
// @Accept json
// @Produce json
// @Param album_id path int true "album_id"
// @Param photo_id path int true "photo_id"
// @Success 204
// @Router /album/upload/{album_id}/{photo_id} [post]
// @Security ApiKeyAuth
func UploadAlbumCover(c *gin.Context) {
	albumParam := c.Param("album_id")
	photoParam := c.Param("photo_id")

	albumID, err := strconv.Atoi(albumParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Error{Message: "ID has to a  be integer: " + err.Error()})
		return
	}
	photoID, err := strconv.Atoi(photoParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Error{Message: "ID has to a  be integer: " + err.Error()})
		return
	}

	paramID, found := c.Get("id")
	if !found {
		c.Status(http.StatusBadRequest)
	}

	tokenUserID := paramID.(uint64)

	err = albumUseCase.UploadAlbumCover(
		tokenUserID,
		uint64(albumID),
		uint64(photoID),
		photoRepo.GetPhoto,
		albumRepo.GetByID,
		albumRepo.Update,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &dto.Error{Message: "error updating album cover: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, http.StatusNoContent)

}

// @BasePath /v1

// Update - receives id as a parameter, update the album, and returns an album objects
// @Summary - Update album
// @Description - Update an album
// @Tags - Album
// @Accept json
// @Produce json
// @Param Album body dto.Album true "Album to be updated"
// @Success 200 {object} dto.Album
// @Router /album [put]
// @Security ApiKeyAuth
func Update(c *gin.Context) {

	albumDTO := &dto.Album{}

	err := c.ShouldBindJSON(&albumDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Error{Message: "Cannot bind album to json: " + err.Error()})
		return
	}

	paramID, found := c.Get("id")
	if !found {
		c.Status(http.StatusBadRequest)
		return
	}

	userID := paramID.(uint64)

	albumVO := albumDTO.ConvertToVO()

	err = albumUseCase.Update(userID, albumVO, albumRepo.Update)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Error{Message: "Cannot update an album: " + err.Error()})
		return
	}

	updatedAlbumDTO := &dto.Album{}
	updatedAlbumDTO.ParseFromVO(albumVO)

	c.JSON(http.StatusOK, updatedAlbumDTO)
}

// @BasePath /v1

// Delete - receives id as a parameter, and returns a string
// @Summary - Delete album
// @Schemes - $ref: '#/definitions/dto.Album'
// @Description - Delete an album
// @Tags - Album
// @Accept application/x-json-stream
// @Produce json
// @Param id path int true "id"
// @Error 500 {object} dto.Error
// @Success 204
// @Router /album/{id} [delete]
// @Security ApiKeyAuth
func Delete(c *gin.Context) {
	albumID := c.Param("id")
	intAlbumID, err := strconv.Atoi(albumID)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Error{Message: "ID has to be a integer: " + err.Error()})
		return
	}

	paramID, found := c.Get("id")
	if !found {
		c.Status(http.StatusBadRequest)
		return
	}

	userID := paramID.(uint64)

	err = albumUseCase.Delete(userID, uint64(intAlbumID), albumRepo.Delete, albumRepo.GetByID)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Error{Message: "failed in to delete a user: " + err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
