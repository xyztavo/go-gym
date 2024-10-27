package database

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/xyztavo/go-gym/internal/models"
)

func CreateGymExerciseRoutineCollection(gymId string, gymExerciseCollection *models.CreateGymExerciseRoutineCollection) (createdGymCollectionId string, err error) {
	id, _ := gonanoid.New()
	if err = db.QueryRow("INSERT INTO gyms_exercises_routines_collections (id, name, description, gym_id, exercise_routine_collection_id) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		id, gymExerciseCollection.Name, gymExerciseCollection.Description, gymId, gymExerciseCollection.ExerciseRoutineCollectionId).Scan(&createdGymCollectionId); err != nil {
		return "", err
	}
	return createdGymCollectionId, nil
}
