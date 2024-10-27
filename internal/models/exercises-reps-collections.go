package models

type CreateExerciseRepsCollection struct {
	CollectionId  string `json:"collectionId" validate:"required"`
	ExerciseRepId string `json:"exerciseRepId" validate:"required"`
}

type ExerciseRoutineCollectionRoutine struct {
	Id                       string `json:"id"`
	ExerciseRepsCollectionId string `json:"exerciseRepsCollectionId"`
	ExerciseRepsId           string `json:"exerciseRepsId"`
}
