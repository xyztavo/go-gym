package database

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/xyztavo/go-gym/internal/models"
)

func CreateExercise(exercise *models.CreateExercise) (createdExerciseId string, err error) {
	id, _ := gonanoid.New()
	err = db.QueryRow("INSERT INTO exercises (id, name, description, gif) VALUES ($1, $2, $3, $4) RETURNING id",
		id, exercise.Name, exercise.Description, exercise.Gif).Scan(&createdExerciseId)
	if err != nil {
		return "", err
	}
	return createdExerciseId, nil
}
