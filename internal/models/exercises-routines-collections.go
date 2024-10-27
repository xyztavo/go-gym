package models

type CreateExerciseRoutineColletion struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type ExerciseRoutineColletion struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateExerciseRoutineCollectionRoutine struct {
	ExerciseRoutineCollectionId string `json:"exerciseRoutineCollectionId" validate:"required"`
	RoutineId                   string `json:"exerciseRoutineId" validate:"required"`
}

type ExerciseRoutineCollectionRoutine struct {
	Id                         string `json:"id"`
	ExerciseRoutineColletionId string `json:"exerciseRoutineCollectionId"`
	ExerciseRoutineId          string `json:"exerciseRoutineId"`
}
