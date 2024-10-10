package database

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/xyztavo/go-gym/internal/models"
)

func CreateExerciseRoutine(exerciseRoutine *models.CreateExerciseRoutine) (createdExerciseRoutineId string, err error) {
	id, _ := gonanoid.New()
	err = db.QueryRow("INSERT INTO exercises_routines (id, routine_id, exercise_id, reps, sets) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		id, exerciseRoutine.RoutineId, exerciseRoutine.ExerciseId, exerciseRoutine.Reps, exerciseRoutine.Sets).Scan(&createdExerciseRoutineId)
	if err != nil {
		return "", err
	}
	return createdExerciseRoutineId, nil
}
