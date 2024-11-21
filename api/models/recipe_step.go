package models

import (
	"github.com/ZayenJS/database"
	"github.com/ZayenJS/dto"
)

type RecipeStep struct {
	StepId    int     `json:"step_id"`
	RecipeId  int     `json:"recipe_id"`
	Text      string  `json:"text"`
	Order     int     `json:"order"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
}

func RecipeStepTableName() string {
	return "recipe_steps"
}

func RecipeStepFromDTO(recipeId int, stepDto *dto.StepDTO) *RecipeStep {
	return &RecipeStep{
		RecipeId: recipeId,
		Text:     stepDto.Text,
		Order:    stepDto.Order,
	}
}

func (recipeStep *RecipeStep) Save() error {
	stmt, err := database.Db.Prepare("INSERT INTO recipe_steps (recipe_id, text, `order`) VALUES (?, ?, ?)")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(recipeStep.RecipeId, recipeStep.Text, recipeStep.Order)

	if err != nil {
		return err
	}

	stmt, err = database.Db.Prepare("SELECT recipe_step_id FROM recipe_steps WHERE recipe_id = ? AND `order` = ?")

	if err != nil {
		return err
	}

	defer stmt.Close()

	err = stmt.QueryRow(recipeStep.RecipeId, recipeStep.Order).Scan(&recipeStep.StepId)

	return err
}
