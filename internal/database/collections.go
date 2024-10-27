package database

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/xyztavo/go-gym/internal/models"
)

func CreateCollection(collection *models.CreateCollection) (createdRoutineCollectionId string, err error) {
	id, _ := gonanoid.New()
	if err = db.QueryRow("INSERT INTO collections (id, name, description) VALUES ($1, $2, $3) RETURNING id", id, collection.Name, collection.Description).Scan(&createdRoutineCollectionId); err != nil {
		return "", err
	}
	return createdRoutineCollectionId, nil
}

func GetCollections() (collections []models.Collection, err error) {
	rows, err := db.Query("SELECT * FROM exercises_routines_collections")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var collection models.Collection
		rows.Scan(&collection.Id, &collection.Name, &collection.Description)
		if err := rows.Err(); err != nil {
			return nil, err
		}
		collections = append(collections, collection)
	}
	return collections, nil
}

func GetCollectionsByRoutineId(routineId *models.GetCollectionsByRoutineId) (collections []models.Collection, err error) {
	rows, err := db.Query(`
	SELECT c.id, c.name, c.description FROM routines AS r 
	JOIN routines_exercises_reps_collections AS rerc ON r.id = rerc.routine_id 
	JOIN exercises_reps_collections AS erc ON rerc.exercise_reps_collection_id = erc.id 
	JOIN collections AS c ON erc.collection_id = c.id WHERE r.id = $1`, routineId.RoutineId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var collection models.Collection
		rows.Scan(&collection.Id, &collection.Name, &collection.Description)
		if err := rows.Err(); err != nil {
			return nil, err
		}
		collections = append(collections, collection)
	}
	return collections, nil
}
