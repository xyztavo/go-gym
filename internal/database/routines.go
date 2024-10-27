package database

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/xyztavo/go-gym/internal/models"
)

func CreateRoutine(routine *models.CreateRoutine) (createdRoutineId string, err error) {
	id, _ := gonanoid.New()
	if err = db.QueryRow("INSERT INTO routines (id, name, description) VALUES ($1, $2, $3) RETURNING id",
		id, routine.Name, routine.Description).
		Scan(&createdRoutineId); err != nil {
		return "", nil
	}
	return createdRoutineId, nil
}

func GetRoutines() (routines []models.Routine, err error) {
	rows, err := db.Query("SELECT * FROM routines")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, err
		}
		var routine models.Routine
		rows.Scan(&routine.Id, &routine.Name, &routine.Description)
		routines = append(routines, routine)
	}
	return routines, nil
}
