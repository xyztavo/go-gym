package models

type CreateCollection struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Img         string `json:"img" validate:"required"`
}

type Collection struct {
	Id          string `json:"id"`
	AdminId     string `json:"adminId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Img         string `json:"img"`
}

type RoutineCollection struct {
	Id                  string `json:"id"`
	RoutineCollectionId string `json:"routineCollectionId"`
	AdminId             string `json:"adminId"`
	Name                string `json:"name"`
	Description         string `json:"description"`
	Img                 string `json:"img"`
}

type GetCollectionsByRoutineId struct {
	RoutineId string `json:"routineId" validate:"required"`
}
