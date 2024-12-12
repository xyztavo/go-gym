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

func GetCollections(query string, page int) ([]models.Collection, int, error) {
	res := 20
	pageOffset := res * page
	query = "%" + query + "%"
	rows, err := db.Query(`
		WITH TotalCount AS (
			SELECT COUNT(*) AS total FROM collections WHERE name ILIKE $1
		)
		SELECT 
			c.id, 
			c.admin_id, 
			c.name, 
			c.description, 
			c.img, 
			TotalCount.total
		FROM collections c, TotalCount
		WHERE c.name ILIKE $1
		LIMIT $2 OFFSET $3
	`, query, res, pageOffset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()
	var collections []models.Collection
	var total int
	for rows.Next() {
		var collection models.Collection
		if err := rows.Scan(&collection.Id, &collection.AdminId, &collection.Name, &collection.Description, &collection.Img, &total); err != nil {
			return nil, 0, err
		}
		collections = append(collections, collection)
	}
	maxPages := (total + res - 1) / res
	return collections, maxPages, nil
}

func GetCollectionById(id string) (collection models.Collection, err error) {
	if err := db.QueryRow("SELECT * FROM collections WHERE id = $1", id).Scan(&collection.Id, &collection.AdminId, &collection.Name, &collection.Description, &collection.Img); err != nil {
		return collection, err
	}
	return collection, nil
}

func GetCollectionsByRoutineId(routineId string) (collections []models.Collection, err error) {
	rows, err := db.Query(`
	SELECT rc.id, c.admin_id, c.name, c.description, c.img FROM routines_collections AS rc 
		LEFT JOIN collections AS c ON rc.collection_id = c.id WHERE rc.routine_id = $1`, routineId)
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
