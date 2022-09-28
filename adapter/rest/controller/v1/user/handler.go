package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	userRepo "wallet-api/adapter/database/sql/user"
	"wallet-api/internal/dto"
	userUseCase "wallet-api/internal/use_cases/user"
)

// GetByID - receives id as a parameter, and returns a user object
// @Summary - Get one user
// @Description Get user by id
// @Tags - User
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} dto.User
// @Router /user/{id} [get]
// @Security ApiKeyAuth
func GetByID(c *gin.Context) {
	userID := c.Param("id")
	userIDParam, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Error{Message: "ID has to a be integer: " + err.Error()})
		return
	}

	_, found := c.Get("id")
	if !found {
		c.Status(http.StatusBadRequest)
		return
	}

	found, user, err := userUseCase.GetByID(uint64(userIDParam), userRepo.GetByID)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Error{Message: "Cannot find a user: " + err.Error()})
		return
	}

	if !found {
		c.Status(http.StatusNotFound)
		return
	}

	userDTO := &dto.User{}
	userDTO.ParseFromVO(user)

	c.JSON(http.StatusOK, userDTO)
}

// @BasePath /v1

// Create - create a user objects, and returns a user objects
// @Summary - Create user
// @Description - Create a user
// @Tags - User
// @Accept json
// @Produce json
// @Param User body dto.UserCreateRequest true "User to be created"
// @Success 201 {object} dto.UserCreateResponse
// @Router /user [post]
// @Security ApiKeyAuth
func Create(c *gin.Context) {
	userCreateRequestDTO := &dto.UserCreateRequest{}

	err := c.BindJSON(userCreateRequestDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Error{Message: "Cannot bind user to json: " + err.Error()})
		return
	}

	userVO := userCreateRequestDTO.GenerateUserVO()
	user, err := userUseCase.Create(userVO, userRepo.Create, userRepo.GetByEmail, userRepo.GetByLogin)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Error{Message: "Cannot create a user: " + err.Error()})
		return
	}

	userCreateResponseDTO := &dto.UserCreateResponse{}
	userCreateResponseDTO.ParseFromUserVO(user)

	c.JSON(http.StatusCreated, userCreateResponseDTO)
}

// @BasePath /v1

// Upload - upload a profile picture
// @Summary - upload a profile picture
// @Description - upload a profile picture
// @Tags - User
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Param file formData file true  "upload your file"
// @Success 204
// @Router /user/upload/{id} [post]
// @Security ApiKeyAuth
func Upload(c *gin.Context) {
	userID := c.Param("id")
	userIDParam, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Error{Message: "ID has to a be integer: " + err.Error()})
		return
	}

	idParam, found := c.Get("id")
	if !found {
		c.Status(http.StatusBadRequest)
		return
	}
	tokenUserID := idParam.(uint64)

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, &dto.Error{Message: "Cannot get a photo: " + err.Error()})
		return
	}

	err = userUseCase.Upload(tokenUserID, uint64(userIDParam), userRepo.GetByID, userRepo.Update, file)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Error{Message: "error to upload photo: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, http.StatusNoContent)

}

// @BasePath /v1

// Update - receives id as a parameter, update a user, and returns a user objects
// @Summary - Update user
// @Description - Update a user by id
// @Tags - User
// @Accept json
// @Produce json
// @Param User body dto.User true "User to be updated"
// @Success 200 {object} dto.User
// @Router /user/{id} [put]
// @Security ApiKeyAuth
func Update(c *gin.Context) {
	userDTO := &dto.User{}

	err := c.ShouldBindJSON(userDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Error{Message: "Cannot bind to json: " + err.Error()})
		return
	}

	paramID, found := c.Get("id")
	if !found {
		c.Status(http.StatusBadRequest)
		return
	}

	userID := paramID.(uint64)

	userVO := userDTO.ConvertToVO()

	err = userUseCase.Update(userID, userVO, userRepo.Update, userRepo.GetByEmail)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Error{Message: "Cannot update a user: " + err.Error()})
		return
	}

	c.Status(http.StatusNoContent)

}

// @BasePath /v1

// Delete - receives id as a parameter, and returns a string
// @Summary - Delete user
// @Description - Delete a user by id
// @Tags - User
// @Accept json
// @Produce json
// @Success 204
// @Param id path int true "id"
// @Error 500 {object} dto.Error
// @Router /user/{id} [delete]
// @Security ApiKeyAuth
func Delete(c *gin.Context) {
	userId := c.Param("id")
	userIDParam, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Error{Message: "ID has to be a integer: " + err.Error()})
		return
	}

	idParam, found := c.Get("id")
	if !found {
		c.Status(http.StatusBadRequest)
		return
	}

	tokenUserID := idParam.(uint64)

	err = userUseCase.Delete(tokenUserID, uint64(userIDParam), userRepo.Delete)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Error{Message: "failed to delete a user: " + err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
