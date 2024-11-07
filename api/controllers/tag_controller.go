package controllers

import (
	"github.com/ZayenJS/models"
	response "github.com/ZayenJS/utils"
	"github.com/gin-gonic/gin"
)

func GetAllTags(c *gin.Context) {
	httpResponse := response.New(c)

	tags, err := models.GetAllTags()

	if err != nil {
		httpResponse.Error(500, err)
		return
	}

	httpResponse.JSON(200, gin.H{"tags": tags})
}
