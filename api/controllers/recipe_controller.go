package controllers

import (
	"errors"
	"net/http"

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

func SearchRecipesByName(c *gin.Context) {
	name := c.Query("name")
	response := response.New(c)

	if name == "" {
		err := errors.New("name is required")
		response.Error(http.StatusBadRequest, err)
		return
	}

	recipes, err := repository.NewRecipeRepository().SearchByName(name)

	if err != nil {
		response.Error(http.StatusInternalServerError, err)
		return
	}

	if len(recipes) == 0 {
		err = errors.New("no recipes found")
		response.Error(http.StatusNotFound, err)
		return
	}

	response.JSON(http.StatusOK, recipes)
}
