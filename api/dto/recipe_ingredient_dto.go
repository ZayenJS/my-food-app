package dto

type CreateRecipeIngredientDTO struct {
	FoodId   int    `json:"food_id" binding:"required"`
	Quantity int    `json:"quantity" binding:"required"`
	Unit     string `json:"unit" binding:"required"`
}
