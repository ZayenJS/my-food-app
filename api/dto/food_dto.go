package dto

import "strings"

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
}

func (f *CreateFoodDTO) GetErrors() map[string]string {
	var createFoodDtoErrors = make(map[string]string)

	if f.Name == "" {
		createFoodDtoErrors["name"] = "name is required"
	}

	if f.Brand == "" {
		createFoodDtoErrors["brand"] = "brand is required"
	}

	if f.Calories == 0 {
		createFoodDtoErrors["calories"] = "calories is required"
	}

	if f.Protein == 0 {
		createFoodDtoErrors["protein"] = "protein is required"
	}

	if f.Carbs == 0 {
		createFoodDtoErrors["carbs"] = "carbs is required"
	}

	if f.Sugar == 0 {
		createFoodDtoErrors["sugar"] = "sugar is required"
	}

	if f.Fat == 0 {
		createFoodDtoErrors["fat"] = "fat is required"
	}

	if f.SaturatedFat == 0 {
		createFoodDtoErrors["saturated_fat"] = "saturated fat is required"
	}

	if f.Fiber == 0 {
		createFoodDtoErrors["fiber"] = "fiber is required"
	}

	if f.Sodium == 0 {
		createFoodDtoErrors["sodium"] = "sodium is required"
	}

	return createFoodDtoErrors
}

func (f *CreateFoodDTO) NormalizeNames() {
	f.Brand = strings.Trim(strings.ToLower(f.Brand), " ")
	f.Name = strings.Trim(strings.ToLower(f.Name), " ")
}
