package dto

type TagDTO struct {
	Name  string `json:"name" binding:"required"`
	Color string `json:"color"`
}
