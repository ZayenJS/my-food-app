package router

import (
	"github.com/ZayenJS/controllers"
	"github.com/gin-gonic/gin"
)

func setupTagRoutes(engine *gin.RouterGroup) {
	tag := engine.Group("/tag")
	tag.GET("", controllers.GetAllTags)
}
