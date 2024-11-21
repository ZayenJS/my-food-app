package repository

import (
	"github.com/ZayenJS/dto"
	"github.com/ZayenJS/models"
)

type FoodRepository struct {
}

func NewFoodRepository() *FoodRepository {
	return &FoodRepository{}
}

func (r *FoodRepository) Create(foodDto *dto.CreateFoodDTO, brandId int) (*models.Food, error) {
	food := models.FoodFromCreateDTO(foodDto, brandId)

	err := food.Save()

	return food, err
}
