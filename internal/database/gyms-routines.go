package database

import (
	"errors"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/xyztavo/go-gym/internal/models"
)

func CreateGymRoutine(gymId string, routineId string) (createdGymRoutineId string, err error) {
	id, _ := gonanoid.New()
	if err = db.QueryRow("INSERT INTO gyms_routines (id, routine_id, gym_id) VALUES ($1, $2, $3) RETURNING id",
		id, routineId, gymId).
		Scan(&createdGymRoutineId); err != nil {
		return "", err
	}
	return createdGymRoutineId, nil
}

func GetGymRoutines(gymId string) (gymRoutines []models.Routine, err error) {
	rows, err := db.Query("SELECT gr.id, r.name, r.description, r.img FROM gyms_routines AS gr LEFT JOIN routines AS r ON gr.routine_id = r.id WHERE gym_id = $1", gymId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		if err = rows.Err(); err != nil {
			return nil, err
		}
		var gymRoutine models.Routine
		if err = rows.Scan(&gymRoutine.Id, &gymRoutine.Name, &gymRoutine.Description, &gymRoutine.Img); err != nil {
			return nil, err
		}
		gymRoutines = append(gymRoutines, gymRoutine)
	}
	return gymRoutines, nil
}

func DeleteGymRoutine(adminId string, id string) (err error) {
	admin, err := GetUserById(adminId)
	if err != nil {
		return err
	}
	gym, err := GetGymById(*admin.GymId)
	if err != nil {
		return err
	}
	if gym.AdminId != adminId {
		return errors.New("user is not the admin of the gym")
	}
	_, err = db.Exec("DELETE FROM gyms_routines WHERE id = $1", id)
	return err
}
