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
func GetExercises(query string, page int) ([]models.Exercise, int, error) {
	res := 20
	pageOffset := res * page
	query = "%" + query + "%"
	rows, err := db.Query(`
        WITH TotalCount AS (
            SELECT COUNT(*) AS total FROM exercises WHERE name ILIKE $1
        )
        SELECT 
            e.id, 
            e.name, 
            e.description, 
            e.gif, 
            TotalCount.total
        FROM exercises e, TotalCount
        WHERE e.name ILIKE $1
        LIMIT $2 OFFSET $3
    `, query, res, pageOffset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()
	var exercises []models.Exercise
	var total int
	for rows.Next() {
		var exercise models.Exercise
		if err := rows.Scan(&exercise.Id, &exercise.Name, &exercise.Description, &exercise.Gif, &total); err != nil {
			return nil, 0, err
		}
		exercises = append(exercises, exercise)
	}
	maxPages := (total + res - 1) / res
	return exercises, maxPages, nil
}
