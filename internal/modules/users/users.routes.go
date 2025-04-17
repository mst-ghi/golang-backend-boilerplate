package users

import (
	"app/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	ctrl := NewUsersController()

	authGroup := router.Group("/users").Use(middlewares.Auth)
	{
		authGroup.GET("", ctrl.FindAll)
		authGroup.GET("/:id", ctrl.FindOne)
	}
}
