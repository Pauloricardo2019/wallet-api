package router

import (
	loginController "wallet-api/adapter/rest/controller/v1/login"
	photoController "wallet-api/adapter/rest/controller/v1/photo"
	userController "wallet-api/adapter/rest/controller/v1/user"
	"wallet-api/adapter/rest/middleware"

	"github.com/gin-gonic/gin"
)

func InitializeRouter(router *gin.RouterGroup) {
	userGroup := router.Group("user")
	{
		userGroup.GET("/:id", middleware.Auth(), userController.GetByID)
		userGroup.POST("/", userController.Create)
		userGroup.POST("/upload/:id", middleware.Auth(), userController.Upload)
		userGroup.PUT("/:id", middleware.Auth(), userController.Update)
		userGroup.DELETE("/:id", middleware.Auth(), userController.Delete)
	}

	loginGroup := router.Group("auth")
	{
		loginGroup.POST("/", loginController.Login)
	}

	photoGroup := router.Group("photo", middleware.Auth())
	{
		photoGroup.POST("/:album_id", photoController.Upload)
		photoGroup.GET("/:album_id", photoController.GetAllPhotosByAlbum)
		photoGroup.GET("/album/:photo_id", photoController.GetPhoto)
		photoGroup.DELETE("/:photo_id", photoController.DeletePhoto)
	}

}
