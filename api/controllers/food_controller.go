package controllers

import (
	"net/http"
	"strconv"

	"github.com/ZayenJS/appHttp"
	"github.com/ZayenJS/models"
	"github.com/ZayenJS/utils"
	"github.com/gin-gonic/gin"
)

func GetAllFoods(c *gin.Context) {
	httpResponse := appHttp.NewResponse(c)

	page := c.DefaultQuery("page", "1")
	perPage := c.DefaultQuery("per-page", "10")
	sort := utils.ValidateSqlSortType(c.DefaultQuery("sort", "asc"))
	indexBy := c.Query("index-by")

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
		foods, err := models.GetFoodsByName(query, convertedPerPage, (convertedPage-1)*convertedPerPage, sort)

		if err != nil {
			httpResponse.Error(500, err)
			return
		}

		if indexBy == "letter" {
			httpResponse.JSON(http.StatusOK, utils.IndexFoodsByLetter(foods))
			return
		}

		httpResponse.JSON(200, foods)
		return
	}

	foods, err := models.GetAllFoods(convertedPerPage, (convertedPage-1)*convertedPerPage, sort)

	if err != nil {
		httpResponse.Error(500, err)
		return
	}

	if indexBy == "letter" {
		httpResponse.JSON(http.StatusOK, utils.IndexFoodsByLetter(foods))
		return
	}

	httpResponse.JSON(200, foods)
}
