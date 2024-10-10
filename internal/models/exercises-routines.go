package models

type ExerciseRoutine struct {
	Id         string `json:"id"`
	RoutineId  string `json:"routineId"`
	ExerciseId string `json:"exerciseId"`
	Reps       string `json:"reps"`
}

type CreateExerciseRoutine struct {
	RoutineId  string `json:"routineId" validate:"required"`
	ExerciseId string `json:"exerciseId" validate:"required"`
	Reps       int    `json:"reps" validate:"required"`
	Sets       int    `json:"sets" validate:"required"`
}
