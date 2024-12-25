package database

import (
	"errors"

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

func GetRoutines(query string, page int) ([]models.Routine, int, error) {
	res := 20
	pageOffset := res * page
	query = "%" + query + "%"
	rows, err := db.Query(`
		WITH TotalCount AS (
			SELECT COUNT(*) AS total FROM routines WHERE name ILIKE $1
		)
		SELECT 
			r.id, 
			r.admin_id, 
			r.name, 
			r.description, 
			r.img, 
			TotalCount.total
		FROM routines r, TotalCount
		WHERE r.name ILIKE $1
		LIMIT $2 OFFSET $3
	`, query, res, pageOffset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()
	var routines []models.Routine
	var total int
	for rows.Next() {
		var routine models.Routine
		if err := rows.Scan(&routine.Id, &routine.AdminId, &routine.Name, &routine.Description, &routine.Img, &total); err != nil {
			return nil, 0, err
		}
		routines = append(routines, routine)
	}
	maxPages := (total + res - 1) / res
	return routines, maxPages, nil
}

func GetUserRoutines(adminId string) (routines []models.Routine, err error) {
	rows, err := db.Query("SELECT * FROM routines WHERE admin_id = $1", adminId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var routine models.Routine
		if err := rows.Scan(&routine.Id, &routine.AdminId, &routine.Name, &routine.Description, &routine.Img); err != nil {
			return nil, err
		}
		routines = append(routines, routine)
	}
	return routines, nil
}

func UpdateRoutine(id string, adminId string, routine *models.UpdateRoutine) (err error) {
	routineById, err := GetRoutineById(id)
	if err != nil {
		return err
	}
	if adminId != routineById.AdminId {
		return errors.New("user is not routine admin")
	}
	_, err = db.Exec("UPDATE routines SET name = $1, description = $2, img = $3 WHERE id = $4", routine.Name, routine.Description, routine.Img, id)
	return err
}

func DeleteRoutine(id string, adminId string) (err error) {
	routineById, err := GetRoutineById(id)
	if err != nil {
		return err
	}
	if adminId != routineById.AdminId {
		return errors.New("user is not routine admin")
	}
	_, err = db.Exec("DELETE FROM routines WHERE id = $1", id)
	return err
}
