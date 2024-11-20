package database

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/xyztavo/go-gym/internal/models"
)

func CreateCollection(adminId string, collection *models.CreateCollection) (createdRoutineCollectionId string, err error) {
	id, _ := gonanoid.New()
	if err = db.QueryRow("INSERT INTO collections (id, admin_id, name, description, img) VALUES ($1, $2, $3, $4, $5) RETURNING id", id, adminId, collection.Name, collection.Description, collection.Img).Scan(&createdRoutineCollectionId); err != nil {
		return "", err
	}
	return createdRoutineCollectionId, nil
}

func GetCollections() (collections []models.Collection, err error) {
	rows, err := db.Query("SELECT * FROM collections")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var collection models.Collection
		rows.Scan(&collection.Id, &collection.AdminId, &collection.Name, &collection.Description, &collection.Img)
		if err := rows.Err(); err != nil {
			return nil, err
		}
		collections = append(collections, collection)
	}
	return collections, nil
}

func GetCollectionById(id string) (collection models.Collection, err error) {
	if err := db.QueryRow("SELECT * FROM collections WHERE id = $1", id).Scan(&collection.Id, &collection.AdminId, &collection.Name, &collection.Description, &collection.Img); err != nil {
		return collection, err
	}
	return collection, nil
}

func GetCollectionsByRoutineId(routineId *models.GetCollectionsByRoutineId) (collections []models.Collection, err error) {
	rows, err := db.Query(`
	SELECT c.id, c.name, c.description, c.img FROM routines AS r 
	JOIN routines_exercises_reps_collections AS rerc ON r.id = rerc.routine_id 
	JOIN exercises_reps_collections AS erc ON rerc.exercise_reps_collection_id = erc.id 
	JOIN collections AS c ON erc.collection_id = c.id WHERE r.id = $1`, routineId.RoutineId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var collection models.Collection
		rows.Scan(&collection.Id, &collection.Name, &collection.Description, &collection.Img)
		if err := rows.Err(); err != nil {
			return nil, err
		}
		collections = append(collections, collection)
	}
	return collections, nil
}

func GetAdminCollections(adminId string) (collections []models.Collection, err error) {
	rows, err := db.Query("SELECT * FROM collections WHERE admin_id = $1", adminId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var collection models.Collection
		rows.Scan(&collection.Id, &collection.AdminId, &collection.Name, &collection.Description, &collection.Img)
		if err := rows.Err(); err != nil {
			return nil, err
		}
		collections = append(collections, collection)
	}
	return collections, nil
}
