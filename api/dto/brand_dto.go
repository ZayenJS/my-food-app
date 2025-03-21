package dto

type CreateBrandDTO struct {
	Name string `json:"name" binding:"required"`
}
