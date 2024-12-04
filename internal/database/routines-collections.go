package database

import (
	"errors"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

func CreateRoutineCollection(adminId string, routineId string, collectionId string) (createdRoutineExerciseRepsCollectionId string, err error) {
	id, _ := gonanoid.New()
	routine, err := GetRoutineById(routineId)
	if err != nil {
		return "", err
	}
	if adminId != routine.AdminId {
		return "", errors.New("user is not routine admin")
	}

	if err = db.QueryRow("INSERT INTO routines_collections (id, admin_id, routine_id, collection_id) VALUES ($1, $2, $3, $4) RETURNING id",
		id, adminId, routineId, collectionId).Scan(&createdRoutineExerciseRepsCollectionId); err != nil {
		return "", err
	}
	return createdRoutineExerciseRepsCollectionId, nil
}
