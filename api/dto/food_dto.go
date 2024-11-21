package dto

type CreateFoodDTO struct {
	Name         string  `json:"name" binding:"required"`
	Brand        string  `json:"brand" binding:"required"`
	Calories     float64 `json:"calories" binding:"required"`
	Protein      float64 `json:"protein" binding:"required"`
	Carbs        float64 `json:"carbs" binding:"required"`
	Sugar        float64 `json:"sugar" binding:"required"`
	Fat          float64 `json:"fat" binding:"required"`
	SaturatedFat float64 `json:"saturated_fat" binding:"required"`
	Fiber        float64 `json:"fiber" binding:"required"`
	Sodium       float64 `json:"sodium" binding:"required"`
	ImageUrl     string  `json:"image_url"`
	Category     string  `json:"category" binding:"required"`
}
