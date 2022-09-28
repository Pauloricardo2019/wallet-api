package comment

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	albumRepo "wallet-api/adapter/database/sql/album"
	commentRepo "wallet-api/adapter/database/sql/comment"
	photoRepo "wallet-api/adapter/database/sql/photo"
	"wallet-api/adapter/rest/utils"
	"wallet-api/internal/dto"
	commentUseCase "wallet-api/internal/use_cases/comment"
)

// @BasePath /v1

// Create - receive an objects
// @Summary - create a comments
// @Description - create a comments
// @Tags - Comment
// @Accept json
// @Produce json
// @Param Comment body dto.CommentRequest true "comment data"
// @Success 201
// @Router /comment [post]
// @Security ApiKeyAuth
func Create(c *gin.Context) {

	commentDTO := &dto.CommentRequest{}

	idParam, found := c.Get("id")
	if !found {
		c.Status(http.StatusBadRequest)
		return
	}

	userID := idParam.(uint64)

	err := c.BindJSON(&commentDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Error{Message: "Cannot bind object to json: " + err.Error()})
		return
	}

	commentVO := commentDTO.ConvertToVO()

	if commentVO.PhotoID != nil {
		found, _, err = commentUseCase.SaveCommentPhoto(
			userID,
			commentVO,
			photoRepo.GetPhoto,
			commentRepo.CreateComment,
			commentRepo.IncrementCommentCountPhoto,
		)
		if err != nil {
			c.JSON(http.StatusBadRequest, &dto.Error{Message: "can't comments this photo: " + err.Error()})
			return
		}
		if !found {
			c.JSON(http.StatusNotFound, &dto.Error{Message: "not found: " + err.Error()})
			return
		}
	}

	if commentVO.AlbumID != nil {
		found, _, err = commentUseCase.SaveCommentAlbum(
			userID,
			commentVO,
			albumRepo.GetByID,
			commentRepo.CreateComment,
			commentRepo.IncrementCommentCountAlbum,
		)
		if err != nil {
			c.JSON(http.StatusBadRequest, &dto.Error{Message: "can't comments this album: " + err.Error()})
			return
		}
		if !found {
			c.JSON(http.StatusNotFound, &dto.Error{Message: "not found: " + err.Error()})
			return
		}

	}

	c.JSON(http.StatusCreated, nil)

}

// @BasePath /v1

// GetAllPhotosComment - receive a photo_id , and show all comments
// @Summary - get all comments
// @Description - get all comments
// @Tags - Comment
// @Accept json
// @Produce json
// @Param photo_id path int true "photo_id"
// @Param   limit     query     int     true  "limit per page"
// @Param   page      query     int     true  "number of page"
// @Param   sort      query     string  false  "search sort"
// @Success 200 {array} dto.Comment
// @Router /comment/{photo_id} [get]
// @Security ApiKeyAuth
func GetAllPhotosComment(c *gin.Context) {

	pagination := utils.GeneratePaginationFromRequest(c)

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

	found, comments, err := commentUseCase.GetAllCommentsByPhoto(&pagination, uint64(photoID), commentRepo.GetCommentByPhotoID)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Error{Message: "it was not possible to get all the comments: " + err.Error()})
		return
	}
	if !found {
		c.Status(http.StatusNotFound)
		return
	}

	commentsDTO := &dto.Comment{}
	commentsArrayDTO := commentsDTO.ParseFromArrayVO(comments)

	c.JSON(http.StatusOK, &commentsArrayDTO)
}

// @BasePath /v1

// GetAllAlbumsComment - receive an album_id , and show all comments
// @Summary - get all comments
// @Description - get all comments
// @Tags - Comment
// @Accept json
// @Produce json
// @Param album_id path int true "album_id"
// @Param   limit     query     int     true  "limit per page"
// @Param   page      query     int     true  "number of page"
// @Param   sort      query     string  false  "search sort"
// @Success 200 {array} dto.Comment
// @Router /comment/album/{album_id} [get]
// @Security ApiKeyAuth
func GetAllAlbumsComment(c *gin.Context) {

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

	found, comments, err := commentUseCase.GetAllCommentsByAlbum(&pagination, uint64(albumID), commentRepo.GetCommentByAlbumID)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Error{Message: "it was not possible to get all the comments: " + err.Error()})
		return
	}
	if !found {
		c.Status(http.StatusNotFound)
		return
	}

	commentsDTO := &dto.Comment{}
	commentsArrayDTO := commentsDTO.ParseFromArrayVO(comments)

	c.JSON(http.StatusOK, &commentsArrayDTO)
}

// @BasePath /v1

// Delete - receive a comment id , and delete this comment
// @Summary - delete a comments
// @Description - delete a comments
// @Tags - Comment
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 204
// @Router /comment/{id} [delete]
// @Security ApiKeyAuth
func Delete(c *gin.Context) {

	commentIDParam := c.Param("id")
	commentID, err := strconv.Atoi(commentIDParam)
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

	err = commentUseCase.DeleteComment(
		userID,
		uint64(commentID),
		commentRepo.GetComment,
		commentRepo.DeleteComment,
		commentRepo.DecrementCommentCountPhoto,
		commentRepo.DecrementCommentCountAlbum,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Error{Message: "Can't delete this comment: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}
