package database

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/xyztavo/go-gym/internal/models"
)

func CreateRoutine(routine *models.CreateRoutine) (createdRoutineId string, err error) {
	id, _ := gonanoid.New()
	err = db.QueryRow("INSERT INTO routines (id, name, description, thumb) VALUES ($1, $2, $3, $4) RETURNING id",
		id, routine.Name, routine.Description, routine.Thumb).Scan(&createdRoutineId)
	if err != nil {
		return "", err
	}
	return createdRoutineId, nil
}

func GetRoutines() (routines []models.Routine, err error) {
	rows, err := db.Query("SELECT * FROM routines")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var routine models.Routine
		rows.Scan(&routine.Id, &routine.Name, &routine.Description, &routine.Thumb)
		routines = append(routines, routine)
		if err := rows.Err(); err != nil {
			return nil, err
		}
	}
	return routines, nil
}
