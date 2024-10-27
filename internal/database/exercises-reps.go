package database

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/xyztavo/go-gym/internal/models"
)

func CreateExerciseReps(exerciseReps *models.CreateExerciseReps) (createdExerciseRepsId string, err error) {
	id, _ := gonanoid.New()
	err = db.QueryRow("INSERT INTO exercises_reps (id, exercise_id, reps, sets) VALUES ($1, $2, $3, $4) RETURNING id",
		id, exerciseReps.ExerciseId, exerciseReps.Reps, exerciseReps.Sets).Scan(&createdExerciseRepsId)
	if err != nil {
		return "", err
	}
	return createdExerciseRepsId, nil
}
