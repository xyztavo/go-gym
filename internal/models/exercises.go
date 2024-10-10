package models

type Exercise struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Gif         string `json:"gif"`
}
type CreateExercise struct {
	Id          string `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Gif         string `json:"gif" validate:"required"`
}
