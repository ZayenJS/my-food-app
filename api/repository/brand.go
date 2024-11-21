package repository

import (
	"github.com/ZayenJS/database"
	"github.com/ZayenJS/models"
)

type BrandRepository struct {
}

func NewBrandRepository() *BrandRepository {
	return &BrandRepository{}
}

func (r *BrandRepository) GetBrandByName(brandName string) (*models.Brand, error) {
	stms, err := database.Db.Prepare("SELECT * FROM brands WHERE name = ?")

	if err != nil {
		return nil, err
	}

	row := stms.QueryRow(brandName)
	brand := models.Brand{}
	err = row.Scan(&brand.BrandId, &brand.Name, &brand.CreatedAt, &brand.UpdatedAt)

	return &brand, err
}
