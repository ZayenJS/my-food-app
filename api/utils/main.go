package utils

import "github.com/ZayenJS/models"

func ValidateSqlSortType(sort string) string {
	if sort != "asc" && sort != "desc" {
		return "asc"
	}

	return sort
}

func IndexFoodsByLetter(foods []models.Food) map[string][]models.Food {
	foodsIndex := make(map[string][]models.Food)

	for _, food := range foods {
		firstLetter := string(food.Name[0])
		foodsIndex[firstLetter] = append(foodsIndex[firstLetter], food)
	}

	return foodsIndex
}
