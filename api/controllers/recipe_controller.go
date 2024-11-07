package controllers

import (
	"net/http"
	"time"

	"github.com/ZayenJS/dto"
	"github.com/ZayenJS/models"
	"github.com/gin-gonic/gin"
)

func CreateRecipe(c *gin.Context) {
	var recipeDto dto.RecipeDTO
	if err := c.ShouldBindJSON(&recipeDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	now := time.Now().Format("2000-01-01 00:00:00")

	recipe := models.RecipeFromDTO(recipeDto)
	recipe.CreatedAt = now
	recipe.UpdatedAt = nil

	recipe.Save()

	for _, ingredientDto := range recipeDto.Ingredients {
		recipeIngredient := models.RecipeIngredientFromDTO(recipe.Id, ingredientDto)
		recipeIngredient.CreatedAt = now
		recipeIngredient.Save()
	}

	for _, stepDto := range recipeDto.Steps {
		recipeStep := models.RecipeStepFromDTO(recipe.Id, stepDto)
		recipeStep.CreatedAt = now
		recipeStep.Save()
	}

	for _, tagId := range recipeDto.Tags {
		recipeTag := models.RecipeTagFromDTO(recipe.Id, tagId)
		recipeTag.CreatedAt = now
		recipeTag.Save()
	}

	c.JSON(http.StatusCreated, gin.H{"status": "created"})
}
