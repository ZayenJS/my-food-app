package dto

type CreateRecipeDTO struct {
	Name        string                      `json:"name" binding:"required"`
	Description string                      `json:"description" binding:"required"`
	Difficulty  int                         `json:"difficulty" binding:"required"`
	PrepTime    int                         `json:"prep_time" binding:"required"`
	CookTime    int                         `json:"cook_time" binding:"required"`
	RestTime    int                         `json:"rest_time"`
	Servings    int                         `json:"servings" binding:"required"`
	ImageUrl    string                      `json:"image_url"`
	Ingredients []CreateRecipeIngredientDTO `json:"ingredients" binding:"required"`
	Steps       []StepDTO                   `json:"steps" binding:"required"`
	Tags        []int                       `json:"tags" binding:"required"`
}

type RecipeForUIDTO struct {
	Id          int                        `json:"id"`
	Name        string                     `json:"name"`
	Description string                     `json:"description"`
	Difficulty  int                        `json:"difficulty"`
	PrepTime    int                        `json:"prep_time"`
	CookTime    int                        `json:"cook_time"`
	RestTime    *int                       `json:"rest_time"`
	Servings    int                        `json:"servings"`
	ImageUrl    string                     `json:"image_url"`
	Ingredients []RecipeIngredientForUIDTO `json:"ingredients"`
	Steps       []RecipeStepForUIDTO       `json:"steps"`
	Tags        []RecipeTagForUIDTO        `json:"tags"`
	Macros      RecipeMacrosDTO            `json:"macros"`
	CreatedAt   string                     `json:"created_at"`
	UpdatedAt   *string                    `json:"updated_at"`
}

type RecipeIngredientForUIDTO struct {
	CreateRecipeIngredientDTO
	Name string `json:"name"`
}

type RecipeStepForUIDTO struct {
	Id int `json:"id"`
	StepDTO
}

type RecipeTagForUIDTO struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type RecipeMacrosDTO struct {
	CaloriesPer100g  float64 `json:"calories_per_100g"`
	ProteinPer100g   float64 `json:"protein_per_100g"`
	CarbsPer100g     float64 `json:"carbs_per_100g"`
	SugarPer100g     float64 `json:"sugar_per_100g"`
	FatPer100g       float64 `json:"fat_per_100g"`
	SaturatedPer100g float64 `json:"saturated_per_100g"`
	FiberPer100g     float64 `json:"fiber_per_100g"`
	SodiumPer100g    float64 `json:"sodium_per_100g"`
}
