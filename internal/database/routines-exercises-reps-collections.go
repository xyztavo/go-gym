package database

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/xyztavo/go-gym/internal/models"
)

func CreateRoutineExerciseRepsCollection(createRoutineExerciseRepsCollection *models.CreateRoutineExerciseRepsCollection) (createdRoutineExerciseRepsCollectionId string, err error) {
	id, _ := gonanoid.New()
	if err = db.QueryRow("INSERT INTO routines_exercises_reps_collections (id, routine_id, exercise_reps_collection_id) VALUES ($1, $2, $3) RETURNING id",
		id, createRoutineExerciseRepsCollection.RoutineId, createRoutineExerciseRepsCollection.ExerciseRepsCollectionId).Scan(&createdRoutineExerciseRepsCollectionId); err != nil {
		return "", err
	}
	return createdRoutineExerciseRepsCollectionId, nil
}
