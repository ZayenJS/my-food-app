package router

import (
	"github.com/ZayenJS/controllers"
	"github.com/gin-gonic/gin"
)

func setupFoodRoutes(engine *gin.RouterGroup) {
	food := engine.Group("/food")
	food.GET("", controllers.GetAllFoods)
	food.POST("", controllers.CreateFood)
}
