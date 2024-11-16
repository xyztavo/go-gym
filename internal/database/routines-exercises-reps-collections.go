package database

import (
	"errors"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/xyztavo/go-gym/internal/models"
)

func CreateRoutineExerciseRepsCollection(adminId string, createRoutineExerciseRepsCollection *models.CreateRoutineExerciseRepsCollection) (createdRoutineExerciseRepsCollectionId string, err error) {
	id, _ := gonanoid.New()
	routine, err := GetRoutineById(createRoutineExerciseRepsCollection.RoutineId)
	if err != nil {
		return "", err
	}
	if adminId != routine.AdminId {
		return "", errors.New("user is not routine admin")
	}

	if err = db.QueryRow("INSERT INTO routines_exercises_reps_collections (id, admin_id, routine_id, exercise_reps_collection_id) VALUES ($1, $2, $3, $4) RETURNING id",
		id, adminId, createRoutineExerciseRepsCollection.RoutineId, createRoutineExerciseRepsCollection.ExerciseRepsCollectionId).Scan(&createdRoutineExerciseRepsCollectionId); err != nil {
		return "", err
	}
	return createdRoutineExerciseRepsCollectionId, nil
}
