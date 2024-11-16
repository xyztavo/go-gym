package database

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/xyztavo/go-gym/internal/models"
)

func CreateRoutine(adminId string, routine *models.CreateRoutine) (createdRoutineId string, err error) {
	id, _ := gonanoid.New()
	if err = db.QueryRow("INSERT INTO routines (id, admin_id, name, description, img) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		id, adminId, routine.Name, routine.Description, routine.Img).
		Scan(&createdRoutineId); err != nil {
		return "", err
	}
	return createdRoutineId, nil
}

func GetRoutineById(id string) (routine models.Routine, err error) {
	if err := db.QueryRow("SELECT * FROM routines WHERE id = $1", id).Scan(&routine.Id, &routine.AdminId, &routine.Name, &routine.Description, &routine.Img); err != nil {
		return routine, err
	}
	return routine, nil
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
		rows.Scan(&routine.Id, &routine.AdminId, &routine.Name, &routine.Description, &routine.Img)
		routines = append(routines, routine)
	}
	return routines, nil
}
