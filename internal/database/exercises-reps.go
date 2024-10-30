package database

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/xyztavo/go-gym/internal/models"
)

func CreateExerciseReps(exerciseReps *models.CreateExerciseReps) (createdExerciseRepsId string, err error) {
	id, _ := gonanoid.New()
	err = db.QueryRow("INSERT INTO exercises_reps (id, exercise_id, reps, sets) VALUES ($1, $2, $3, $4) RETURNING id",
		id, exerciseReps.ExerciseId, exerciseReps.Reps, exerciseReps.Sets).Scan(&createdExerciseRepsId)
	if err != nil {
		return "", err
	}
	return createdExerciseRepsId, nil
}

func GetExercisesRepsByCollectionId(getExerciseRepsByCollectionId *models.GetExercisesRepsByCollectionId) (exercisesReps []models.ExerciseRepsWithName, err error) {
	rows, err := db.Query(`
	SELECT e.name, e.description, er.reps, er.sets FROM collections AS c 
		JOIN exercises_reps_collections AS erc ON c.id = erc.collection_id
		JOIN exercises_reps AS er ON er.id = erc.exercise_rep_id
		JOIN exercises AS e ON er.exercise_id = e.id WHERE c.id = $1
	`, getExerciseRepsByCollectionId.CollectionId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var exerciseReps models.ExerciseRepsWithName
		if err = rows.Scan(&exerciseReps.Name, &exerciseReps.Description, &exerciseReps.Reps, &exerciseReps.Sets); err != nil {
			return nil, err
		}
		exercisesReps = append(exercisesReps, exerciseReps)
	}
	return exercisesReps, nil
}
