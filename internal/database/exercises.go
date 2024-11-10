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

func GetExercises(query string) (exercises []models.Exercise, err error) {
	query = "%" + query + "%"
	rows, err := db.Query("SELECT * FROM exercises WHERE name LIKE $1", query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var exercise models.Exercise
		rows.Scan(&exercise.Id, &exercise.Name, &exercise.Description, &exercise.Gif)
		exercises = append(exercises, exercise)
		if err := rows.Err(); err != nil {
			return nil, err
		}
	}
	return exercises, nil
}
