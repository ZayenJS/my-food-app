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
	RestTime    int     `json:"rest_time"`
	Servings    int     `json:"servings"`
	Rating      *int    `json:"rating"`
	ImageUrl    string  `json:"image_url"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   *string `json:"updated_at"`
}

type RecipeWithRelations struct {
	Recipe
	Ingredients []RecipeIngredient `json:"ingredients"`
	Steps       []RecipeStep       `json:"steps"`
	Tags        []Tag              `json:"tags"`
}

func RecipeTableName() string {
	return "recipes"
}

func RecipeFromCreateDTO(recipeDto *dto.CreateRecipeDTO) *Recipe {
	return &Recipe{
		Name:        recipeDto.Name,
		Description: recipeDto.Description,
		Difficulty:  recipeDto.Difficulty,
		PrepTime:    recipeDto.PrepTime,
		CookTime:    recipeDto.CookTime,
		RestTime:    recipeDto.RestTime,
		Servings:    recipeDto.Servings,
		ImageUrl:    recipeDto.ImageUrl,
	}
}

func (recipe *Recipe) Save() error {
	stmt, err := database.Db.Prepare("INSERT INTO recipes (name, description, difficulty, prep_time, cook_time, rest_time, servings, image_url) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(recipe.Name, recipe.Description, recipe.Difficulty, recipe.PrepTime, recipe.CookTime, recipe.RestTime, recipe.Servings, recipe.ImageUrl)

	if err != nil {
		return err
	}

	stmt, err = database.Db.Prepare("SELECT id FROM recipes WHERE name = ? AND description = ?")

	if err != nil {
		return err
	}

	defer stmt.Close()

	err = stmt.QueryRow(recipe.Name, recipe.Description).Scan(&recipe.Id)

	return err
}
