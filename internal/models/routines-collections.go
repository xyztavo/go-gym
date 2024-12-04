package models

type CreateRoutineCollection struct {
	CollectionId string `json:"collectionId" validate:"required"`
}
