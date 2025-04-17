package app

import (
	"app/internal"
	"app/internal/modules/auth"
	"app/internal/modules/users"

	"github.com/gin-gonic/gin"
)

// @tags    App
// @router	/api [get]
// @summary	app route, get heathy status
func HomeRoute(c *gin.Context) {
	ctrl := internal.GetController()
	ctrl.Success(c, map[string]string{
		"heathy": "I'm OK...",
	})
}

func RegisterRoutes(router *gin.RouterGroup) {
	router.GET("", HomeRoute)

	v1Group := router.Group("/v1")
	{
		auth.RegisterRoutes(v1Group)
		users.RegisterRoutes(v1Group)
	}
}
