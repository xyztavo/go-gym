package database

import (
	"errors"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/xyztavo/go-gym/internal/models"
)

func CreateExercisesRepCollection(adminId string, addExerciseRoutinesCollectionRoutines *models.CreateExerciseRepsCollection) (createdaddRoutinesCollectionRoutinesId string, err error) {
	id, _ := gonanoid.New()
	collection, err := GetCollectionById(addExerciseRoutinesCollectionRoutines.CollectionId)
	if err != nil {
		return "", err
	}
	if adminId != collection.AdminId {
		return "", errors.New("user is not the collection admin")
	}
	if err = db.QueryRow("INSERT INTO exercises_reps_collections (id, admin_id, collection_id, exercise_id, reps, sets) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		id, adminId, addExerciseRoutinesCollectionRoutines.CollectionId, addExerciseRoutinesCollectionRoutines.ExerciseId,
		addExerciseRoutinesCollectionRoutines.Reps, addExerciseRoutinesCollectionRoutines.Sets).
		Scan(&createdaddRoutinesCollectionRoutinesId); err != nil {
		return "", err
	}
	return createdaddRoutinesCollectionRoutinesId, nil
}

func GetExercisesRepsCollection() (exercisesRoutinesCollectionsRoutines []models.ExerciseRepsCollection, err error) {
	rows, err := db.Query("SELECT * FROM exercises_reps_collections")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var exerciseRoutineCollectionRoutine models.ExerciseRepsCollection
		if err := rows.Err(); err != nil {
			return nil, err
		}
		rows.Scan(&exerciseRoutineCollectionRoutine.Id, &exerciseRoutineCollectionRoutine.ExerciseRepsCollectionId, &exerciseRoutineCollectionRoutine.ExerciseId, &exerciseRoutineCollectionRoutine.Reps, &exerciseRoutineCollectionRoutine.Sets)
		exercisesRoutinesCollectionsRoutines = append(exercisesRoutinesCollectionsRoutines, exerciseRoutineCollectionRoutine)
	}
	return exercisesRoutinesCollectionsRoutines, nil
}

func GetExercisesRepsCollectionsByCollectionId(collectionId *models.GetExercisesRepsByCollectionId) (exercisesRepsCollections []models.ExerciseRepCollectionFormatted, err error) {
	rows, err := db.Query(`
	SELECT e.name, e.description, e.gif, erc.reps, erc.sets 
		FROM collections AS c 
		LEFT JOIN exercises_reps_collections AS erc 
		LEFT JOIN exercises AS e ON erc.exercise_id = e.id
		ON c.id = erc.collection_id WHERE c.id = $1`, collectionId.CollectionId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var exerciseRepsCollection models.ExerciseRepCollectionFormatted
		if err := rows.Scan(&exerciseRepsCollection.Name, &exerciseRepsCollection.Description, &exerciseRepsCollection.Gif,
			&exerciseRepsCollection.Reps, &exerciseRepsCollection.Sets); err != nil {
			return nil, err
		}
		exercisesRepsCollections = append(exercisesRepsCollections, exerciseRepsCollection)
	}
	return exercisesRepsCollections, nil
}
