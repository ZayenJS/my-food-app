package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/ZayenJS/dto"
	"github.com/ZayenJS/repository"
	response "github.com/ZayenJS/utils"
	"github.com/gin-gonic/gin"
)

func CreateRecipe(c *gin.Context) {
	var recipeDto dto.CreateRecipeDTO
	response := response.New(c)
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
	httpResponse := response.New(c)

	page := c.DefaultQuery("page", "1")
	perPage := c.DefaultQuery("per-page", "10")
	name := c.Query("name")

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

	if name != "" {
		recipes, err := repository.NewRecipeRepository().SearchByName(name, convertedPerPage, (convertedPage-1)*convertedPerPage)

		if err != nil {
			httpResponse.Error(500, err)
			return
		}

		httpResponse.JSON(200, gin.H{"recipes": recipes})
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
	response := response.New(c)
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
