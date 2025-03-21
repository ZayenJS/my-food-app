package router

import (
	"github.com/ZayenJS/controllers"
	"github.com/gin-gonic/gin"
)

func setupBrandRoutes(engine *gin.RouterGroup) {
	food := engine.Group("/brand")
	food.GET("", controllers.GetAllBrands)
}
