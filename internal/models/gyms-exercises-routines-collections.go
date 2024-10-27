package models

type CreateGymExerciseRoutineCollection struct {
	Name                        string `json:"name" validate:"required"`
	Description                 string `json:"description" validate:"required"`
	ExerciseRoutineCollectionId string `json:"exerciseRoutineCollectionId" validate:"required"`
}
