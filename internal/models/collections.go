package models

type CreateCollection struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type Collection struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
