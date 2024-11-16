package router

import (
	"github.com/ZayenJS/controllers"
	"github.com/gin-gonic/gin"
)

func setupRecipeRoutes(engine *gin.RouterGroup) {
	recipe := engine.Group("/recipe")
	recipe.POST("", controllers.CreateRecipe)
	recipe.GET("", controllers.GetAllRecipes)
	recipe.POST(":id/rate", controllers.RateRecipe)
}
