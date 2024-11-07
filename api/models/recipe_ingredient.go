package models

import (
	"github.com/ZayenJS/database"
	"github.com/ZayenJS/dto"
)

type RecipeIngredient struct {
	Id        int     `json:"id"`
	RecipeId  int     `json:"recipe_id"`
	FoodId    int     `json:"food_id"`
	Quantity  int     `json:"quantity"`
	Unit      string  `json:"unit"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
}

func RecipeIngredientTableName() string {
	return "recipe_ingredients"
}

func RecipeIngredientFromDTO(recipeId int, ingredientDto dto.RecipeIngredientDTO) *RecipeIngredient {
	return &RecipeIngredient{
		RecipeId: recipeId,
		FoodId:   ingredientDto.FoodId,
		Quantity: ingredientDto.Quantity,
		Unit:     ingredientDto.Unit,
	}
}

func (recipeIngredient *RecipeIngredient) Save() (*RecipeIngredient, error) {
	stmt, err := database.Db.Prepare("INSERT INTO recipe_ingredients (recipe_id, food_id, quantity, unit) VALUES (?, ?, ?, ?)")

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(recipeIngredient.RecipeId, recipeIngredient.FoodId, recipeIngredient.Quantity, recipeIngredient.Unit)

	if err != nil {
		return nil, err
	}

	stmt, err = database.Db.Prepare("SELECT id FROM recipe_ingredients WHERE recipe_id = ? AND food_id = ?")

	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(recipeIngredient.RecipeId, recipeIngredient.FoodId).Scan(&recipeIngredient.Id)

	return recipeIngredient, err
}