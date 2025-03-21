package repository

import (
	"fmt"
	"strings"

	"github.com/ZayenJS/database"
	"github.com/ZayenJS/dto"
	"github.com/ZayenJS/models"
)

type FoodRepository struct {
}

func NewFoodRepository() *FoodRepository {
	return &FoodRepository{}
}

func (r *FoodRepository) GetAllFoods(limit int, offset int, sort string) ([]models.Food, error) {
	stmt, err := database.Db.Prepare(fmt.Sprintf("SELECT * FROM `%v` ORDER BY name %v LIMIT ? OFFSET ?", models.FoodTableName(), sort))

	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(limit, offset)

	if err != nil {
		return nil, err
	}

	foods := []models.Food{}

	for rows.Next() {
		var food models.Food
		err := rows.Scan(&food.FoodId, &food.Name, &food.BrandId, &food.Calories, &food.Protein, &food.Carbs, &food.Sugar, &food.Fat, &food.SaturatedFat, &food.Fiber, &food.Sodium, &food.ImageUrl, &food.CreatedAt, &food.UpdatedAt)

		if err != nil {
			return nil, err
		}

		foods = append(foods, food)
	}

	return foods, nil
}

func (r *FoodRepository) GetFoodsByName(name string, limit int, offset int, sort string) ([]models.Food, error) {
	stmt, err := database.Db.Prepare(fmt.Sprintf("SELECT * FROM `%v` WHERE `name` LIKE ? ORDER BY name %v LIMIT ? OFFSET ?", models.FoodTableName(), sort))

	if err != nil {
		return nil, err
	}

	formattedName := fmt.Sprintf("%%%s%%", strings.ToLower(name))
	rows, err := stmt.Query(formattedName, limit, offset)

	if err != nil {
		return nil, err
	}

	foods := []models.Food{}

	for rows.Next() {
		var food models.Food
		err := rows.Scan(&food.FoodId, &food.Name, &food.BrandId, &food.Calories, &food.Protein, &food.Carbs, &food.Sugar, &food.Fat, &food.SaturatedFat, &food.Fiber, &food.Sodium, &food.ImageUrl, &food.CreatedAt, &food.UpdatedAt)

		if err != nil {
			return nil, err
		}

		foods = append(foods, food)
	}

	return foods, nil
}

func (r *FoodRepository) Create(foodDto *dto.CreateFoodDTO, brandId int) (*models.Food, error) {
	food := models.FoodFromCreateDTO(foodDto, brandId)

	err := food.Save()

	return food, err
}

func (r *FoodRepository) GetFoodByName(name string) (*models.Food, error) {
	food := &models.Food{}
	foods, err := r.GetFoodsByName(name, 1, 0, "asc")

	if err != nil {
		return nil, err
	}

	if len(foods) == 0 {
		return nil, nil
	}

	food = &foods[0]

	return food, err
}

func (r *FoodRepository) GetFoodByNameAndBrand(name string, brandName string) (*models.Food, error) {
	food := &models.Food{}
	brandRepository := NewBrandRepository()
	brand, err := brandRepository.GetBrandByName(brandName)

	if err != nil {
		return nil, err
	}

	if brand == nil {
		return nil, nil
	}

	stmt, err := database.Db.Prepare(fmt.Sprintf(`
		SELECT
			f.food_id,
			f.name,
			f.brand_id,
			f.calories,
			f.protein,
			f.carbs,
			f.sugar,
			f.fat,
			f.saturated_fat,
			f.fiber,
			f.sodium,
			f.image_url,
			f.created_at,
			f.updated_at
		FROM %v f
		LEFT JOIN brands b USING(brand_id)
		WHERE f.name LIKE ?
		AND b.name LIKE ?
		ORDER BY f.name ASC
		`, models.FoodTableName()))

	if err != nil {
		return nil, err
	}

	formattedFoodName := fmt.Sprintf("%%%s%%", strings.ToLower(name))
	formattedBrandName := fmt.Sprintf("%%%s%%", strings.ToLower(brandName))
	row := stmt.QueryRow(formattedFoodName, formattedBrandName)

	err = row.Scan(&food.FoodId, &food.Name, &food.BrandId, &food.Calories, &food.Protein, &food.Carbs, &food.Sugar, &food.Fat, &food.SaturatedFat, &food.Fiber, &food.Sodium, &food.ImageUrl, &food.CreatedAt, &food.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return food, nil
}
