package models

type CreateRoutine struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Img         string `json:"img" validate:"required"`
}

type Routine struct {
	Id          string `json:"id"`
	AdminId     string `json:"adminId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Img         string `json:"img" `
}

type UpdateRoutine struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Img         string `json:"img" validate:"required"`
}
