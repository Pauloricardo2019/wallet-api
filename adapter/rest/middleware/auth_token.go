package middleware

import (
	"net/http"
	tokenRepo "wallet-api/adapter/database/sql/token"
	tokenUseCase "wallet-api/internal/use_cases/token"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		const bearerSchema = "Bearer "

		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token := header[len(bearerSchema):]

		if header != (bearerSchema + token) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		found, id := tokenUseCase.ValidateToken(token, tokenRepo.GetTokenByValue)

		if !found {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("id", id)

	}
}
