package database

import (
	"errors"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/xyztavo/go-gym/internal/models"
)

func CreateExercise(adminId string, exercise *models.CreateExercise) (createdExerciseId string, err error) {
	id, _ := gonanoid.New()
	err = db.QueryRow("INSERT INTO exercises (id, admin_id, name, description, gif) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		id, adminId, exercise.Name, exercise.Description, exercise.Gif).Scan(&createdExerciseId)
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
	maxPages := total / res
	return exercises, maxPages, nil
}

func GetUserExercises(userId string) ([]models.Exercise, error) {
	rows, err := db.Query("SELECT id, admin_id, name, description, gif FROM exercises WHERE admin_id = $1", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var exercises []models.Exercise
	for rows.Next() {
		var exercise models.Exercise
		if err := rows.Scan(&exercise.Id, &exercise.AdminId, &exercise.Name, &exercise.Description, &exercise.Gif); err != nil {
			return nil, err
		}
		exercises = append(exercises, exercise)
	}
	return exercises, nil
}

func GetExerciseById(id string) (models.Exercise, error) {
	var exercise models.Exercise
	err := db.QueryRow("SELECT id, admin_id, name, description, gif FROM exercises WHERE id = $1", id).Scan(&exercise.Id, &exercise.AdminId, &exercise.Name, &exercise.Description, &exercise.Gif)
	if err != nil {
		return exercise, err
	}
	return exercise, nil
}

func UpdateExercise(adminId string, id string, exercise *models.UpdateExercise) error {
	exerciseById, err := GetExerciseById(id)
	if err != nil {
		return err
	}
	if adminId != exerciseById.AdminId {
		return errors.New("user is not the exercise admin")
	}

	_, err = db.Exec("UPDATE exercises SET name = $1, description = $2, gif = $3 WHERE id = $4", exercise.Name, exercise.Description, exercise.Gif, id)
	return err
}

func DeleteExercise(adminId string, id string) error {
	exerciseById, err := GetExerciseById(id)
	if err != nil {
		return err
	}
	if adminId != exerciseById.AdminId {
		return errors.New("user is not the exercise admin")
	}
	_, err = db.Exec("DELETE FROM exercises WHERE id = $1", id)
	return err
}
