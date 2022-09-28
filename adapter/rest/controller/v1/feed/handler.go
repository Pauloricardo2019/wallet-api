package feed

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /v1

// GetFeed - check authorization
// @Summary - check authorization
// @Description - feed
// @Tags - Feed
// @Accept json
// @Produce json
// @Success 200
// @Router /feed [get]
// @Security ApiKeyAuth
func GetFeed(c *gin.Context) {

	c.JSON(http.StatusOK, nil)

}
