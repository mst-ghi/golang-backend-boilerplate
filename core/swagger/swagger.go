package swagger

import (
	"app/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterSwagger(router *gin.RouterGroup) {
	docs.SwaggerInfo.Title = "Your app name"
	docs.SwaggerInfo.Description = "Your swagger description"
	docs.SwaggerInfo.Version = "0.1.0"

	router.GET(
		"/docs/*any",
		ginSwagger.WrapHandler(
			swaggerFiles.Handler,
			ginSwagger.DefaultModelsExpandDepth(-1),
		),
	)
}
