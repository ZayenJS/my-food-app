package models

import (
	"fmt"
	"strings"

	"github.com/ZayenJS/database"
)

type Food struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	BrandId      int     `json:"brandId"`
	Calories     int     `json:"calories"`
	Protein      int     `json:"protein"`
	Carbs        int     `json:"carbs"`
	Sugar        int     `json:"sugar"`
	Fat          int     `json:"fat"`
	SaturatedFat int     `json:"saturatedFat"`
	Fiber        int     `json:"fiber"`
	Sodium       int     `json:"sodium"`
	ImageUrl     string  `json:"imageUrl"`
	Category     string  `json:"category"`
	CreatedAt    string  `json:"createdAt"`
	UpdatedAt    *string `json:"updatedAt"`
}

func FoodTableName() string {
	return "foods"
}

func GetAllFoods(limit int, offset int) ([]Food, error) {
	stmt, err := database.Db.Prepare(fmt.Sprintf("SELECT * FROM `%v` LIMIT ? OFFSET ?", FoodTableName()))

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
		err := rows.Scan(&food.Id, &food.Name, &food.BrandId, &food.Calories, &food.Protein, &food.Carbs, &food.Sugar, &food.Fat, &food.SaturatedFat, &food.Fiber, &food.Sodium, &food.ImageUrl, &food.Category, &food.CreatedAt, &food.UpdatedAt)

		if err != nil {
			return nil, err
		}

		foods = append(foods, food)
	}

	return foods, nil
}

func GetFoodsByName(name string, limit int, offset int) ([]Food, error) {
	stmt, err := database.Db.Prepare(fmt.Sprintf("SELECT * FROM `%v` WHERE `name` LIKE ? LIMIT ? OFFSET ?", FoodTableName()))

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
		err := rows.Scan(&food.Id, &food.Name, &food.BrandId, &food.Calories, &food.Protein, &food.Carbs, &food.Sugar, &food.Fat, &food.SaturatedFat, &food.Fiber, &food.Sodium, &food.ImageUrl, &food.Category, &food.CreatedAt, &food.UpdatedAt)

		if err != nil {
			return nil, err
		}

		foods = append(foods, food)
	}

	return foods, nil
}
