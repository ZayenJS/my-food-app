package dto

type StepDTO struct {
	Text  string `json:"text" binding:"required"`
	Order int    `json:"order" binding:"required"`
}
