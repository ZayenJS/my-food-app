package models

import (
	"github.com/ZayenJS/database"
	"github.com/ZayenJS/dto"
)

type Recipe struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Difficulty  int     `json:"difficulty"`
	PrepTime    int     `json:"prep_time"`
	CookTime    int     `json:"cook_time"`
	Servings    int     `json:"servings"`
	ImageUrl    string  `json:"image_url"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   *string `json:"updated_at"`
}

func RecipeTableName() string {
	return "recipes"
}

func RecipeFromDTO(recipeDto dto.RecipeDTO) *Recipe {
	return &Recipe{
		Name:        recipeDto.Name,
		Description: recipeDto.Description,
		Difficulty:  recipeDto.Difficulty,
		PrepTime:    recipeDto.PrepTime,
		CookTime:    recipeDto.CookTime,
		Servings:    recipeDto.Servings,
		ImageUrl:    recipeDto.ImageUrl,
	}
}

func (recipe *Recipe) Save() (*Recipe, error) {
	stmt, err := database.Db.Prepare("INSERT INTO recipes (name, description, difficulty, prep_time, cook_time, servings, image_url) VALUES (?, ?, ?, ?, ?, ?, ?)")

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(recipe.Name, recipe.Description, recipe.Difficulty, recipe.PrepTime, recipe.CookTime, recipe.Servings, recipe.ImageUrl)

	if err != nil {
		return nil, err
	}

	stmt, err = database.Db.Prepare("SELECT id FROM recipes WHERE name = ? AND description = ?")

	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(recipe.Name, recipe.Description).Scan(&recipe.Id)

	return recipe, err
}
