package models

type Plan struct {
	Id          string  `json:"id"`
	GymId       string  `json:"gymId"`
	Name        string  `json:"name" `
	Description string  `json:"description" `
	Price       float32 `json:"price" `
	Duration    int     `json:"duration" `
}

type CreatePlan struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float32 `json:"price" validate:"required"`
	Duration    int     `json:"duration" validate:"required"`
}
