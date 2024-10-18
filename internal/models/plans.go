package models

type CreatePlan struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float32 `json:"price" validate:"required"`
	Duration    int     `json:"duration" validate:"required"`
}
