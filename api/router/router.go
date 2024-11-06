package router

import (
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func Setup() *gin.Engine {
	isDev := os.Getenv("ENV") == "development"

	if isDev {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.Default()
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	setupFoodRoutes(engine)

	return engine
}
