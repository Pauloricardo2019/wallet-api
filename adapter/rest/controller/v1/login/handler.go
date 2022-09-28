package login

import (
	"github.com/gin-gonic/gin"
	"net/http"
	tokenRepo "wallet-api/adapter/database/sql/token"
	userRepo "wallet-api/adapter/database/sql/user"
	"wallet-api/internal/dto"
	UserUseCase "wallet-api/internal/use_cases/user"
)

// @BasePath /v1

// Login - receive an objects loginRequest , and returns a Value objects
// @Summary - Authenticate user
// @Description - Authenticate the user giving a token of Authorization
// @Tags - Login
// @Accept json
// @Produce json
// @Param User body dto.LoginRequest true "User to be connected"
// @Success 200 {object} dto.Token
// @Router /auth [post]
func Login(c *gin.Context) {
	loginRequest := &dto.LoginRequest{}

	err := c.BindJSON(loginRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Error{Message: "Cannot bind object to json: " + err.Error()})
		return
	}

	loginVO := loginRequest.ConvertToVO()
	found, token, err := UserUseCase.Login(loginVO, tokenRepo.CreateToken, userRepo.GetByEmail)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Error{Message: "Cannot connect a user: " + err.Error()})
		return
	}

	if !found {
		c.Status(http.StatusNotFound)
		return
	}

	tokenDTO := dto.Token{}
	tokenDTO.ParseFromVO(token)

	c.JSON(http.StatusOK, tokenDTO)
}
