package controllers

import (
	"strconv"

	food "github.com/ZayenJS/models"
	response "github.com/ZayenJS/utils"
	"github.com/gin-gonic/gin"
)

func GetAllFoods(c *gin.Context) {
	httpResponse := response.New(c)

	page := c.DefaultQuery("page", "1")
	perPage := c.DefaultQuery("per-page", "10")

	convertedPage, err := strconv.Atoi(page)
	if err != nil {
		httpResponse.Error(500, err)
		return
	}
	convertedPerPage, err := strconv.Atoi(perPage)
	if err != nil {
		httpResponse.Error(500, err)
		return
	}

	query := c.Query("name")

	if query != "" {
		foods, err := food.GetByName(query, convertedPerPage, (convertedPage-1)*convertedPerPage)

		if err != nil {
			httpResponse.Error(500, err)
			return
		}

		httpResponse.JSON(200, gin.H{"foods": foods})
		return
	}

	foods, err := food.GetAll(convertedPerPage, (convertedPage-1)*convertedPerPage)

	if err != nil {
		httpResponse.Error(500, err)
		return
	}

	httpResponse.JSON(200, gin.H{"foods": foods})
}
