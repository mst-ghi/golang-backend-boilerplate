package auth

import (
	"app/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	ctrl := NewAuthController()

	guestGroup := router.Group("/auth")
	{
		guestGroup.POST("/login", ctrl.Login)
		guestGroup.POST("/register", ctrl.Register)
		guestGroup.POST("/refresh", ctrl.Refresh)
	}

	authGroup := router.Group("/auth").Use(middlewares.Auth)
	{
		authGroup.GET("/user", ctrl.User)
		authGroup.GET("/logout", ctrl.Logout)
		authGroup.PUT("/password", ctrl.ChangePassword)
	}
}
