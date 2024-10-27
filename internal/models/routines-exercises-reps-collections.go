package models

type CreateRoutineExerciseRepsCollection struct {
	RoutineId                string `json:"routineId" validate:"required"`
	ExerciseRepsCollectionId string `json:"exerciseRepsCollectionId" validate:"required"`
}
