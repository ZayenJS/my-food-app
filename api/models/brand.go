package models

type Brand struct {
	BrandId   int     `json:"brand_id"`
	Name      string  `json:"name"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
}
