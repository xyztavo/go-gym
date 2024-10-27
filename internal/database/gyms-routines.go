package database

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/xyztavo/go-gym/internal/models"
)

func CreateGymRoutine(gymId string, createGymRoutine *models.CreateGymRoutine) (createdGymRoutineId string, err error) {
	id, _ := gonanoid.New()
	if err = db.QueryRow("INSERT INTO gyms_routines (id, name, description, routine_id, gym_id) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		id, createGymRoutine.Name, createGymRoutine.Description, createGymRoutine.RoutineId, gymId).
		Scan(&createdGymRoutineId); err != nil {
		return "", err
	}
	return createdGymRoutineId, nil
}

func GetGymRoutines(gymId string) (gymRoutines []models.GymRoutine, err error) {
	rows, err := db.Query("SELECT * FROM gyms_routines WHERE gym_id = $1", gymId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		if err = rows.Err(); err != nil {
			return nil, err
		}
		var gymRoutine models.GymRoutine
		if err = rows.Scan(&gymRoutine.Id, &gymRoutine.Name, &gymRoutine.Description, &gymRoutine.RoutineId, &gymRoutine.GymId); err != nil {
			return nil, err
		}
		gymRoutines = append(gymRoutines, gymRoutine)
	}
	return gymRoutines, nil
}
