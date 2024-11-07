package dto

type RecipeDTO struct {
	Name        string                `json:"name" binding:"required"`
	Description string                `json:"description" binding:"required"`
	Difficulty  int                   `json:"difficulty" binding:"required"`
	PrepTime    int                   `json:"prep_time" binding:"required"`
	CookTime    int                   `json:"cook_time" binding:"required"`
	Servings    int                   `json:"servings" binding:"required"`
	ImageUrl    string                `json:"image_url"`
	Ingredients []RecipeIngredientDTO `json:"ingredients" binding:"required"`
	Steps       []StepDTO             `json:"steps" binding:"required"`
	Tags        []int                 `json:"tags" binding:"required"`
}
