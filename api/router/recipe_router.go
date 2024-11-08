package router

import (
	"github.com/ZayenJS/controllers"
	"github.com/gin-gonic/gin"
)

func setupRecipeRoutes(engine *gin.Engine) {
	recipe := engine.Group("/recipe")
	recipe.POST("/", controllers.CreateRecipe)
	recipe.GET("/", controllers.SearchRecipesByName)
}
