package controllers

import (
	"github.com/ZayenJS/appHttp"
	"github.com/ZayenJS/models"
	"github.com/gin-gonic/gin"
)

func GetAllTags(c *gin.Context) {
	httpResponse := appHttp.NewResponse(c)

	tags, err := models.GetAllTags()

	if err != nil {
		httpResponse.Error(500, err)
		return
	}

	httpResponse.JSON(200, tags)
}
