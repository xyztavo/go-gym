package database

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/xyztavo/go-gym/internal/models"
)

func CreateExercisesRoutineCollection(exerciseRoutineColletion *models.CreateExerciseRoutineColletion) (createdRoutineCollectionId string, err error) {
	id, _ := gonanoid.New()
	if err = db.QueryRow("INSERT INTO exercises_routines_collections (id, name, description) VALUES ($1, $2, $3) RETURNING id", id, exerciseRoutineColletion.Name, exerciseRoutineColletion.Description).Scan(&createdRoutineCollectionId); err != nil {
		return "", err
	}
	return createdRoutineCollectionId, nil
}

func GetExercisesRoutinesCollections() (exercisesRoutinesCollection []models.ExerciseRoutineColletion, err error) {
	rows, err := db.Query("SELECT * FROM exercises_routines_collections")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var exerciseRoutineCollection models.ExerciseRoutineColletion
		rows.Scan(&exerciseRoutineCollection.Id, &exerciseRoutineCollection.Name, &exerciseRoutineCollection.Description)
		if err := rows.Err(); err != nil {
			return nil, err
		}
		exercisesRoutinesCollection = append(exercisesRoutinesCollection, exerciseRoutineCollection)
	}
	return exercisesRoutinesCollection, nil
}

func CreateExercisesRoutineCollectionRoutine(addExerciseRoutinesCollectionRoutines models.CreateExerciseRoutineCollectionRoutine) (createdaddRoutinesCollectionRoutinesId string, err error) {
	id, _ := gonanoid.New()
	if err = db.QueryRow("INSERT INTO exercises_routines_collections_routines (id, exercises_routines_collections_id, exercises_routines_id) VALUES ($1, $2, $3) RETURNING id",
		id, addExerciseRoutinesCollectionRoutines.ExerciseRoutineCollectionId, addExerciseRoutinesCollectionRoutines.RoutineId).Scan(&createdaddRoutinesCollectionRoutinesId); err != nil {
		return "", err
	}
	return createdaddRoutinesCollectionRoutinesId, nil
}

func GetExercisesRoutinesCollectionsRoutines() (exercisesRoutinesCollectionsRoutines []models.ExerciseRoutineCollectionRoutine, err error) {
	rows, err := db.Query("SELECT * FROM exercises_routines_collections_routines")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var exerciseRoutineCollectionRoutine models.ExerciseRoutineCollectionRoutine
		if err := rows.Err(); err != nil {
			return nil, err
		}
		rows.Scan(&exerciseRoutineCollectionRoutine.Id, &exerciseRoutineCollectionRoutine.ExerciseRoutineColletionId, &exerciseRoutineCollectionRoutine.ExerciseRoutineId)
		exercisesRoutinesCollectionsRoutines = append(exercisesRoutinesCollectionsRoutines, exerciseRoutineCollectionRoutine)
	}
	return exercisesRoutinesCollectionsRoutines, nil
}
