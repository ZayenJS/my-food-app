package models

import (
	"github.com/ZayenJS/database"
	"github.com/ZayenJS/dto"
)

type RecipeStep struct {
	Id        int     `json:"id"`
	RecipeId  int     `json:"recipe_id"`
	Text      string  `json:"text"`
	Order     int     `json:"order"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
}

func RecipeStepTableName() string {
	return "recipe_steps"
}

func RecipeStepFromDTO(recipeId int, stepDto dto.StepDTO) *RecipeStep {
	return &RecipeStep{
		RecipeId: recipeId,
		Text:     stepDto.Text,
		Order:    stepDto.Order,
	}
}

func (recipeStep *RecipeStep) Save() (*RecipeStep, error) {
	stmt, err := database.Db.Prepare("INSERT INTO recipe_steps (recipe_id, text, `order`) VALUES (?, ?, ?)")

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(recipeStep.RecipeId, recipeStep.Text, recipeStep.Order)

	if err != nil {
		return nil, err
	}

	stmt, err = database.Db.Prepare("SELECT id FROM recipe_steps WHERE recipe_id = ? AND `order` = ?")

	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(recipeStep.RecipeId, recipeStep.Order).Scan(&recipeStep.Id)

	return recipeStep, err
}
