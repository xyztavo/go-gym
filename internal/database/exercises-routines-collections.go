package database

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/xyztavo/go-gym/internal/models"
)

func CreateExercisesRepCollection(addExerciseRoutinesCollectionRoutines *models.CreateExerciseRepsCollection) (createdaddRoutinesCollectionRoutinesId string, err error) {
	id, _ := gonanoid.New()
	if err = db.QueryRow("INSERT INTO exercises_reps_collections (id, collection_id, exercise_rep_id) VALUES ($1, $2, $3) RETURNING id",
		id, addExerciseRoutinesCollectionRoutines.CollectionId, addExerciseRoutinesCollectionRoutines.ExerciseRepId).Scan(&createdaddRoutinesCollectionRoutinesId); err != nil {
		return "", err
	}
	return createdaddRoutinesCollectionRoutinesId, nil
}

func GetExercisesRoutinesCollectionsRoutines() (exercisesRoutinesCollectionsRoutines []models.ExerciseRoutineCollectionRoutine, err error) {
	rows, err := db.Query("SELECT * FROM exercises_reps_collections")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var exerciseRoutineCollectionRoutine models.ExerciseRoutineCollectionRoutine
		if err := rows.Err(); err != nil {
			return nil, err
		}
		rows.Scan(&exerciseRoutineCollectionRoutine.Id, &exerciseRoutineCollectionRoutine.ExerciseRepsCollectionId, &exerciseRoutineCollectionRoutine.ExerciseRepsId)
		exercisesRoutinesCollectionsRoutines = append(exercisesRoutinesCollectionsRoutines, exerciseRoutineCollectionRoutine)
	}
	return exercisesRoutinesCollectionsRoutines, nil
}
