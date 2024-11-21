package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/ZayenJS/appHttp"
	"github.com/ZayenJS/dto"
	"github.com/ZayenJS/repository"
	"github.com/gin-gonic/gin"
)

func CreateRecipe(c *gin.Context) {
	var recipeDto dto.CreateRecipeDTO
	response := appHttp.NewResponse(c)
	if err := c.ShouldBindJSON(&recipeDto); err != nil {
		response.Error(http.StatusBadRequest, err)
		return
	}

	recipeRepository := repository.NewRecipeRepository()
	err := recipeRepository.Create(&recipeDto)

	if err != nil {
		response.Error(http.StatusInternalServerError, err)
		return
	}

	response.JSON(http.StatusCreated, gin.H{"status": "created"})
}

func GetAllRecipes(c *gin.Context) {
	httpResponse := appHttp.NewResponse(c)

	page := c.DefaultQuery("page", "1")
	perPage := c.DefaultQuery("per-page", "10")
	name := c.Query("name")

	convertedPage, err := strconv.Atoi(page)
	if err != nil {
		httpResponse.Error(http.StatusInternalServerError, err)
		return
	}
	convertedPerPage, err := strconv.Atoi(perPage)
	if err != nil {
		httpResponse.Error(http.StatusInternalServerError, err)
		return
	}

	if name != "" {
		recipes, err := repository.NewRecipeRepository().SearchByName(name, convertedPerPage, (convertedPage-1)*convertedPerPage)

		if err != nil {
			httpResponse.Error(http.StatusInternalServerError, err)
			return
		}

		httpResponse.JSON(http.StatusOK, gin.H{"recipes": recipes})
		return
	}

	recipes, err := repository.NewRecipeRepository().SearchByName(name, convertedPerPage, (convertedPage-1)*convertedPerPage)

	if err != nil {
		httpResponse.Error(http.StatusInternalServerError, err)
		return
	}

	if len(recipes) == 0 {
		err = errors.New("no recipes found")
		httpResponse.Error(http.StatusNotFound, err)
		return
	}

	httpResponse.JSON(http.StatusOK, recipes)
}

func RateRecipe(c *gin.Context) {
	response := appHttp.NewResponse(c)
	recipeId := c.Param("id")

	if recipeId == "" {
		err := errors.New("recipe id is required")
		response.Error(http.StatusBadRequest, err)
		return
	}

	convertedRecipeId, err := strconv.Atoi(recipeId)

	if err != nil {
		response.Error(http.StatusBadRequest, err)
		return
	}

	var recipeRatingDto dto.RateRecipeDTO
	if err := c.Bind(&recipeRatingDto); err != nil {
		response.Error(http.StatusBadRequest, err)
		return
	}

	if err != nil {
		response.Error(http.StatusBadRequest, err)
		return
	}

	recipeRepository := repository.NewRecipeRepository()
	err = recipeRepository.Rate(convertedRecipeId, recipeRatingDto.Rating)

	if err != nil {
		response.Error(http.StatusInternalServerError, err)
		return
	}

	response.JSON(http.StatusOK, gin.H{"status": "ok"})
}
