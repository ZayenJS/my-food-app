package router

import (
	"github.com/ZayenJS/controllers"
	"github.com/gin-gonic/gin"
)

func setupFoodRoutes(engine *gin.Engine) {
	food := engine.Group("/food")
	food.GET("/", controllers.GetAllFoods)
}
