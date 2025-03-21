package repository

import (
	"fmt"
	"strings"
	"time"

	"github.com/ZayenJS/database"
	"github.com/ZayenJS/dto"
	"github.com/ZayenJS/models"
)

type RecipeRepository struct {
}

func NewRecipeRepository() *RecipeRepository {
	return &RecipeRepository{}
}

func (r *RecipeRepository) Create(recipeDto *dto.CreateRecipeDTO) error {
	now := time.Now().Format("2000-01-01 00:00:00")
	recipe := models.RecipeFromCreateDTO(recipeDto)
	recipe.CreatedAt = now
	recipe.UpdatedAt = nil

	err := recipe.Save()

	if err != nil {
		return err
	}

	for _, ingredientDto := range recipeDto.Ingredients {
		recipeIngredient := models.RecipeIngredientFromCreateDTO(recipe.RecipeId, &ingredientDto)
		recipeIngredient.CreatedAt = now
		err = recipeIngredient.Save()

		if err != nil {
			return err
		}
	}

	for _, stepDto := range recipeDto.Steps {
		recipeStep := models.RecipeStepFromDTO(recipe.RecipeId, &stepDto)
		recipeStep.CreatedAt = now
		err = recipeStep.Save()

		if err != nil {
			return err
		}
	}

	for _, tagId := range recipeDto.Tags {
		recipeTag := models.RecipeTagFromIds(recipe.RecipeId, tagId)
		recipeTag.CreatedAt = now
		err = recipeTag.Save()

		if err != nil {
			return err
		}
	}

	return nil
}

func (r *RecipeRepository) GetAll(limit int, offset int) ([]dto.RecipeForUIDTO, error) {
	stmt, err := database.Db.Prepare(`
		SELECT
			r.recipe_id,
			r.name,
			r.description,
			r.difficulty,
			r.prep_time,
			r.cook_time,
			r.rest_time,
			r.servings,
			r.image_url,
			r.created_at,
			r.updated_at
		FROM recipes r
		LIMIT ?
		OFFSET ?;
	`)

	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	recipes := make([]dto.RecipeForUIDTO, 0)

	for rows.Next() {
		recipe := dto.RecipeForUIDTO{}
		err = rows.Scan(&recipe.RecipeId, &recipe.Name, &recipe.Description, &recipe.Difficulty, &recipe.PrepTime, &recipe.CookTime, &recipe.RestTime, &recipe.Servings, &recipe.ImageUrl, &recipe.CreatedAt, &recipe.UpdatedAt)
		if err != nil {
			return nil, err
		}

		recipe.Ingredients = []dto.RecipeIngredientForUIDTO{}
		recipe.Steps = []dto.RecipeStepForUIDTO{}
		recipe.Tags = []dto.RecipeTagForUIDTO{}

		recipeIngredients, err := joinRecipeIngredients(recipe.RecipeId)
		if err != nil {
			return nil, err
		}

		ingredientIds := make([]int, 0)
		for _, ingredient := range recipeIngredients {
			ingredientIds = append(ingredientIds, ingredient.FoodId)
		}

		recipeMacros, err := calculateMacros(ingredientIds)
		if err != nil {
			return nil, err
		}

		recipe.Ingredients = recipeIngredients
		recipe.Macros = *recipeMacros

		recipeSteps, err := joinRecipeSteps(recipe.RecipeId)
		if err != nil {
			return nil, err
		}

		recipe.Steps = recipeSteps

		recipeTags, err := joinRecipeTags(recipe.RecipeId)
		if err != nil {
			return nil, err
		}

		recipe.Tags = recipeTags

		recipes = append(recipes, recipe)
	}

	return recipes, nil
}

func (r *RecipeRepository) SearchByName(name string, convertedPerPage int, offset int) ([]dto.RecipeForUIDTO, error) {
	stmt, err := database.Db.Prepare(`
		SELECT
			r.recipe_id,
			r.name,
			r.description,
			r.difficulty,
			r.prep_time,
			r.cook_time,
			r.rest_time,
			r.servings,
			r.rating,
			r.image_url,
			r.created_at,
			r.updated_at
		FROM recipes r
		WHERE r.name LIKE ?
		LIMIT ?
		OFFSET ?;
	`)

	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query("%"+name+"%", convertedPerPage, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	recipes := make([]dto.RecipeForUIDTO, 0)

	for rows.Next() {
		recipe := dto.RecipeForUIDTO{}
		err = rows.Scan(&recipe.RecipeId, &recipe.Name, &recipe.Description, &recipe.Difficulty, &recipe.PrepTime, &recipe.CookTime, &recipe.RestTime, &recipe.Servings, &recipe.Rating, &recipe.ImageUrl, &recipe.CreatedAt, &recipe.UpdatedAt)
		if err != nil {
			return nil, err
		}

		recipe.Ingredients = []dto.RecipeIngredientForUIDTO{}
		recipe.Steps = []dto.RecipeStepForUIDTO{}
		recipe.Tags = []dto.RecipeTagForUIDTO{}

		recipeIngredients, err := joinRecipeIngredients(recipe.RecipeId)
		if err != nil {
			return nil, err
		}

		ingredientIds := make([]int, 0)
		for _, ingredient := range recipeIngredients {
			ingredientIds = append(ingredientIds, ingredient.FoodId)
		}

		recipeMacros, err := calculateMacros(ingredientIds)
		if err != nil {
			return nil, err
		}

		recipe.Ingredients = recipeIngredients
		recipe.Macros = *recipeMacros

		recipeSteps, err := joinRecipeSteps(recipe.RecipeId)
		if err != nil {
			return nil, err
		}

		recipe.Steps = recipeSteps

		recipeTags, err := joinRecipeTags(recipe.RecipeId)
		if err != nil {
			return nil, err
		}

		recipe.Tags = recipeTags

		recipes = append(recipes, recipe)
	}

	return recipes, nil
}

func (r *RecipeRepository) Rate(recipeId int, rating int) error {
	stmt, err := database.Db.Prepare(`
		UPDATE recipes SET rating = ? WHERE recipe_id = ?;
	`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(rating, recipeId)

	return err
}

func joinRecipeIngredients(recipeId int) ([]dto.RecipeIngredientForUIDTO, error) {
	stmt, err := database.Db.Prepare(`
		SELECT
			f.food_id,
			ri.quantity,
			ri.unit,
			f.name
		FROM recipe_ingredients ri
		JOIN foods f USING(food_id)
		WHERE ri.recipe_id = ?;
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(recipeId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ingredients := make([]dto.RecipeIngredientForUIDTO, 0)

	for rows.Next() {
		ingredient := dto.RecipeIngredientForUIDTO{}
		err = rows.Scan(&ingredient.FoodId, &ingredient.Quantity, &ingredient.Unit, &ingredient.Name)
		if err != nil {
			return nil, err
		}

		ingredients = append(ingredients, ingredient)
	}

	return ingredients, nil
}

func joinRecipeSteps(recipeId int) ([]dto.RecipeStepForUIDTO, error) {
	stmt, err := database.Db.Prepare(`
		SELECT
			rs.recipe_step_id,
			rs.text,
			rs.order
		FROM recipe_steps rs
		WHERE rs.recipe_id = ?;
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(recipeId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	steps := make([]dto.RecipeStepForUIDTO, 0)

	for rows.Next() {
		step := dto.RecipeStepForUIDTO{}
		err = rows.Scan(&step.Id, &step.Text, &step.Order)
		if err != nil {
			return nil, err
		}

		steps = append(steps, step)
	}

	return steps, nil
}

func joinRecipeTags(recipeId int) ([]dto.RecipeTagForUIDTO, error) {
	stmt, err := database.Db.Prepare(`
		SELECT
			t.tag_id,
			t.name,
			t.color
		FROM recipe_tags rt
		JOIN tags t USING(tag_id)
		WHERE rt.recipe_id = ?;
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(recipeId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tags := make([]dto.RecipeTagForUIDTO, 0)

	for rows.Next() {
		tag := dto.RecipeTagForUIDTO{}
		err = rows.Scan(&tag.Id, &tag.Name, &tag.Color)
		if err != nil {
			return nil, err
		}

		tags = append(tags, tag)
	}

	return tags, nil
}

func calculateMacros(ingredientIds []int) (*dto.RecipeMacrosDTO, error) {
	if len(ingredientIds) == 0 {
		return nil, nil
	}

	// Create a slice of placeholders for SQL IN clause (e.g., ?, ?, ?)
	placeholders := make([]string, len(ingredientIds))
	for i := range ingredientIds {
		placeholders[i] = "?"
	}

	// Join the placeholders into a string like "?, ?, ?"
	query := fmt.Sprintf(`
		SELECT
			ROUND((SUM(f.calories) / COUNT(f.food_id)), 2),
			ROUND((SUM(f.protein) / COUNT(f.food_id)), 2),
			ROUND((SUM(f.carbs) / COUNT(f.food_id)), 2),
			ROUND((SUM(f.sugar) / COUNT(f.food_id)), 2),
			ROUND((SUM(f.fat) / COUNT(f.food_id)), 2),
			ROUND((SUM(f.saturated_fat) / COUNT(f.food_id)), 2),
			ROUND((SUM(f.fiber) / COUNT(f.food_id)), 2),
			ROUND((SUM(f.sodium) / COUNT(f.food_id)), 2)
		FROM foods f
		WHERE f.food_id IN (%s);`, strings.Join(placeholders, ","))

	stmt, err := database.Db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Convert ingredientIds to an interface slice for Query
	args := make([]interface{}, len(ingredientIds))
	for i, v := range ingredientIds {
		args[i] = v
	}

	// Execute the query with the correct arguments
	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	macros := dto.RecipeMacrosDTO{}

	for rows.Next() {
		err = rows.Scan(
			&macros.CaloriesPer100g,
			&macros.ProteinPer100g,
			&macros.CarbsPer100g,
			&macros.SugarPer100g,
			&macros.FatPer100g,
			&macros.SaturatedPer100g,
			&macros.FiberPer100g,
			&macros.SodiumPer100g,
		)
		if err != nil {
			return nil, err
		}
	}

	return &macros, nil
}
