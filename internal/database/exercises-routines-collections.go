package database

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/xyztavo/go-gym/internal/models"
)

func CreateExercisesRepCollection(addExerciseRoutinesCollectionRoutines *models.CreateExerciseRepsCollection) (createdaddRoutinesCollectionRoutinesId string, err error) {
	id, _ := gonanoid.New()
	if err = db.QueryRow("INSERT INTO exercises_reps_collections (id, collection_id, exercise_id, reps, sets) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		id, addExerciseRoutinesCollectionRoutines.CollectionId, addExerciseRoutinesCollectionRoutines.ExerciseId,
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
