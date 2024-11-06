package models

type Food struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	BrandId      int    `json:"brandId"`
	Calories     int    `json:"calories"`
	Protein      int    `json:"protein"`
	Carbs        int    `json:"carbs"`
	Sugar        int    `json:"sugar"`
	Fat          int    `json:"fat"`
	SaturatedFat int    `json:"saturatedFat"`
	Fiber        int    `json:"fiber"`
	Sodium       int    `json:"sodium"`
	ImageUrl     string `json:"imageUrl"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
}
