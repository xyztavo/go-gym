package models

type CreateRoutine struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Thumb       string `json:"thumb" validate:"required"`
}

type Routine struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Thumb       string `json:"thumb"`
}
