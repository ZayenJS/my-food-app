package models

import (
	"fmt"
	"strings"

	"github.com/ZayenJS/database"
	"github.com/ZayenJS/dto"
)

type Food struct {
	FoodId       int     `json:"food_id"`
	Name         string  `json:"name"`
	BrandId      int     `json:"brandId"`
	Calories     float64 `json:"calories"`
	Protein      float64 `json:"protein"`
	Carbs        float64 `json:"carbs"`
	Sugar        float64 `json:"sugar"`
	Fat          float64 `json:"fat"`
	SaturatedFat float64 `json:"saturated_fat"`
	Fiber        float64 `json:"fiber"`
	Sodium       float64 `json:"sodium"`
	ImageUrl     string  `json:"imageUrl"`
	Category     string  `json:"category"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    *string `json:"updated_at"`
}

func FoodTableName() string {
	return "foods"
}

func GetAllFoods(limit int, offset int, sort string) ([]Food, error) {
	stmt, err := database.Db.Prepare(fmt.Sprintf("SELECT * FROM `%v` ORDER BY name %v LIMIT ? OFFSET ?", FoodTableName(), sort))

	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(limit, offset)

	if err != nil {
		return nil, err
	}

	foods := []Food{}

	for rows.Next() {
		var food Food
		err := rows.Scan(&food.FoodId, &food.Name, &food.BrandId, &food.Calories, &food.Protein, &food.Carbs, &food.Sugar, &food.Fat, &food.SaturatedFat, &food.Fiber, &food.Sodium, &food.ImageUrl, &food.Category, &food.CreatedAt, &food.UpdatedAt)

		if err != nil {
			return nil, err
		}

		foods = append(foods, food)
	}

	return foods, nil
}

func GetFoodsByName(name string, limit int, offset int, sort string) ([]Food, error) {
	stmt, err := database.Db.Prepare(fmt.Sprintf("SELECT * FROM `%v` WHERE `name` LIKE ? ORDER BY name %v LIMIT ? OFFSET ?", FoodTableName(), sort))

	if err != nil {
		return nil, err
	}

	formattedName := fmt.Sprintf("%%%s%%", strings.ToLower(name))
	rows, err := stmt.Query(formattedName, limit, offset)

	if err != nil {
		return nil, err
	}

	foods := []Food{}

	for rows.Next() {
		var food Food
		err := rows.Scan(&food.FoodId, &food.Name, &food.BrandId, &food.Calories, &food.Protein, &food.Carbs, &food.Sugar, &food.Fat, &food.SaturatedFat, &food.Fiber, &food.Sodium, &food.ImageUrl, &food.Category, &food.CreatedAt, &food.UpdatedAt)

		if err != nil {
			return nil, err
		}

		foods = append(foods, food)
	}

	return foods, nil
}

func FoodFromCreateDTO(foodDto *dto.CreateFoodDTO, brandId int) *Food {
	return &Food{
		Name:         foodDto.Name,
		BrandId:      brandId,
		Calories:     foodDto.Calories,
		Protein:      foodDto.Protein,
		Carbs:        foodDto.Carbs,
		Sugar:        foodDto.Sugar,
		Fat:          foodDto.Fat,
		SaturatedFat: foodDto.SaturatedFat,
		Fiber:        foodDto.Fiber,
		Sodium:       foodDto.Sodium,
		ImageUrl:     foodDto.ImageUrl,
		Category:     foodDto.Category,
	}
}

func (f *Food) Save() error {
	stmt, err := database.Db.Prepare(fmt.Sprintf("INSERT INTO `%v` (name, brand_id, calories, protein, carbs, sugar, fat, saturated_fat, fiber, sodium, image_url, category) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", FoodTableName()))

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(f.Name, f.BrandId, f.Calories, f.Protein, f.Carbs, f.Sugar, f.Fat, f.SaturatedFat, f.Fiber, f.Sodium, f.ImageUrl, f.Category)

	if err != nil {
		return err
	}

	stmt, err = database.Db.Prepare(fmt.Sprintf("SELECT food_id FROM `%v` WHERE name = ? AND brand_id = ?", FoodTableName()))

	if err != nil {
		return err
	}

	defer stmt.Close()

	err = stmt.QueryRow(f.Name, f.BrandId).Scan(&f.FoodId)

	return err

}
