package engine

import (
	"app/core/config"
	"app/core/middlewares"
	"app/core/swagger"
	"app/domain"
	"app/pkg/handlers"

	"github.com/gin-gonic/gin"
)

var engine *gin.Engine

func GetEngine() *gin.Engine {
	return engine
}

func Initialize() {
	gin.SetMode(config.Get("GIN_MODE"))

	engine = gin.Default()
	engine.SetTrustedProxies(nil)
	engine.RedirectTrailingSlash = true
	engine.RedirectFixedPath = true

	engine.Use(middlewares.Cors())
	engine.Use(gin.CustomRecovery(handlers.InternalErrorHandler))
}

func Serve(addr ...string) {
	runAddress := config.GetRunAddress()

	if addr != nil {
		runAddress = addr[0]
	}

	engine.Run(runAddress)
}

func RegisterRoutes() {
	routerGroup := GetEngine().Group("api")

	domain.RegisterRoutes(routerGroup)

	swagger.RegisterSwagger(routerGroup)
}
