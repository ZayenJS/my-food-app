package repository

import (
	"strings"

	"github.com/ZayenJS/database"
	"github.com/ZayenJS/models"
)

type BrandRepository struct {
}

func NewBrandRepository() *BrandRepository {
	return &BrandRepository{}
}

func (r *BrandRepository) Create(brand *models.Brand) (*models.Brand, error) {
	stmt, err := database.Db.Prepare("INSERT INTO brands (name) VALUES (?)")

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(strings.ToLower(brand.Name))

	return brand, err
}

func (r *BrandRepository) GetBrandByName(brandName string) (*models.Brand, error) {
	stmt, err := database.Db.Prepare("SELECT * FROM brands WHERE name = ?")

	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(brandName)
	brand := models.Brand{}
	err = row.Scan(&brand.BrandId, &brand.Name, &brand.CreatedAt, &brand.UpdatedAt)

	return &brand, err
}

func (r *BrandRepository) GetAll() ([]models.Brand, error) {
	stmt, err := database.Db.Prepare("SELECT * FROM brands ORDER BY name ASC")

	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()

	if err != nil {
		return nil, err
	}

	var brands []models.Brand

	for rows.Next() {
		brand := models.Brand{}
		err = rows.Scan(&brand.BrandId, &brand.Name, &brand.CreatedAt, &brand.UpdatedAt)

		if err != nil {
			return nil, err
		}

		brands = append(brands, brand)
	}

	return brands, nil
}
