package models

type Exercise struct {
	Id          string `json:"id"`
	AdminId     string `json:"adminId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Gif         string `json:"gif"`
}
type CreateExercise struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Gif         string `json:"gif" validate:"required"`
}

type UpdateExercise struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Gif         string `json:"gif" validate:"required"`
}
