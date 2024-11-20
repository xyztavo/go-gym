package models

type CreateExerciseRepsCollection struct {
	CollectionId string `json:"collectionId" validate:"required"`
	ExerciseId   string `json:"exerciseId" validate:"required"`
	Reps         int    `json:"Reps" validate:"required"`
	Sets         int    `json:"Sets" validate:"required"`
}

type CreateMultipleExercisesRepCollection struct {
	CollectionId                 string                        `json:"collectionId" validate:"required"`
	CreateExerciseRepsCollection []AddExerciseToRepsCollection `json:"exerciseReps"`
}

type AddExerciseToRepsCollection struct {
	ExerciseId string `json:"exerciseId" validate:"required"`
	Reps       int    `json:"reps" validate:"required"`
	Sets       int    `json:"sets" validate:"required"`
}

type ExerciseRepsCollection struct {
	Id                       string `json:"id"`
	AdminId                  string `json:"adminId"`
	ExerciseRepsCollectionId string `json:"exerciseRepsCollectionId"`
	ExerciseId               string `json:"exerciseId"`
	Reps                     int    `json:"Reps"`
	Sets                     int    `json:"Sets"`
}
type GetExercisesRepsByCollectionId struct {
	CollectionId string `json:"collectionId" validate:"required"`
}
type ExerciseRepCollectionFormatted struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Gif         string `json:"gif"`
	Reps        int    `json:"reps"`
	Sets        int    `json:"sets"`
}
