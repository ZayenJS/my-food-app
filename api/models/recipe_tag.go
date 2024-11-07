package models

import (
	"github.com/ZayenJS/database"
)

type RecipeTag struct {
	Id        int     `json:"id"`
	RecipeId  int     `json:"recipe_id"`
	TagId     int     `json:"tag_id"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
}

func RecipeTagTableName() string {
	return "recipe_tags"
}

func RecipeTagFromDTO(recipeId int, tagId int) *RecipeTag {
	return &RecipeTag{
		RecipeId: recipeId,
		TagId:    tagId,
	}
}

func (recipeTag *RecipeTag) Save() (*RecipeTag, error) {
	stmt, err := database.Db.Prepare("INSERT INTO recipe_tags (recipe_id, tag_id) VALUES (?, ?)")

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(recipeTag.RecipeId, recipeTag.TagId)

	if err != nil {
		return nil, err
	}

	return recipeTag, nil
}
