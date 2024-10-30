package models

type ExerciseReps struct {
	Id         string `json:"id"`
	ExerciseId string `json:"exerciseId"`
	Reps       string `json:"reps"`
}

type GetExercisesRepsByCollectionId struct {
	CollectionId string `json:"collectionId" validate:"required"`
}

type ExerciseRepsWithName struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Reps        int    `json:"reps"`
	Sets        int    `json:"sets"`
}

type CreateExerciseReps struct {
	ExerciseId string `json:"exerciseId" validate:"required"`
	Reps       int    `json:"reps" validate:"required"`
	Sets       int    `json:"sets" validate:"required"`
}
