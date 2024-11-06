package router

import (
	"github.com/ZayenJS/controllers"
	"github.com/gin-gonic/gin"
)

func setupFoodRoutes(r *gin.Engine) {
	r.GET("/food", controllers.GetAllFoods)
}
