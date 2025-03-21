package router

import (
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	env := os.Getenv("ENV")

	if env == "development" || env == "dev" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.Default()
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true,           // Whether or not to allow cookies
		MaxAge:           12 * time.Hour, // Preflight cache duration
	}))

	engine.Static("/public", "./public")

	apiEndpoint := engine.Group("/api")

	setupFoodRoutes(apiEndpoint)
	setupBrandRoutes(apiEndpoint)
	setupRecipeRoutes(apiEndpoint)
	setupTagRoutes(apiEndpoint)

	return engine
}
