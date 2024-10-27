package models

type ExerciseReps struct {
	Id         string `json:"id"`
	ExerciseId string `json:"exerciseId"`
	Reps       string `json:"reps"`
}

type CreateExerciseReps struct {
	ExerciseId string `json:"exerciseId" validate:"required"`
	Reps       int    `json:"reps" validate:"required"`
	Sets       int    `json:"sets" validate:"required"`
}
