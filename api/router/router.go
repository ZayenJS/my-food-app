package router

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	isDev := os.Getenv("ENV") == "development"

	if isDev {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.Default()

	setupFoodRoutes(engine)
	setupRecipeRoutes(engine)
	setupTagRoutes(engine)

	return engine
}
