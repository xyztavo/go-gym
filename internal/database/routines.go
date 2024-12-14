package database

import (
	. "github.com/go-jet/jet/v2/postgres"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/xyztavo/go-gym/dbname/public/model"
	. "github.com/xyztavo/go-gym/dbname/public/table"
	"github.com/xyztavo/go-gym/internal/models"
)

func CreateRoutine(adminId string, routine *models.CreateRoutine) (createdRoutineId string, err error) {
	id, _ := gonanoid.New()
	if err = db.QueryRow("INSERT INTO routines (id, admin_id, name, description, img) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		id, adminId, routine.Name, routine.Description, routine.Img).
		Scan(&createdRoutineId); err != nil {
		return "", err
	}
	return createdRoutineId, nil
}

func GetRoutineById(id string) (routine models.Routine, err error) {
	if err := db.QueryRow("SELECT * FROM routines WHERE id = $1", id).Scan(&routine.Id, &routine.AdminId, &routine.Name, &routine.Description, &routine.Img); err != nil {
		return routine, err
	}
	return routine, nil
}

func GetRoutines(query string, page int) ([]models.Routine, int, error) {
	res := 20
	pageOffset := res * page
	query = "%" + query + "%"
	rows, err := db.Query(`
		WITH TotalCount AS (
			SELECT COUNT(*) AS total FROM routines WHERE name ILIKE $1
		)
		SELECT 
			r.id, 
			r.admin_id, 
			r.name, 
			r.description, 
			r.img, 
			TotalCount.total
		FROM routines r, TotalCount
		WHERE r.name ILIKE $1
		LIMIT $2 OFFSET $3
	`, query, res, pageOffset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()
	var routines []models.Routine
	var total int
	for rows.Next() {
		var routine models.Routine
		if err := rows.Scan(&routine.Id, &routine.AdminId, &routine.Name, &routine.Description, &routine.Img, &total); err != nil {
			return nil, 0, err
		}
		routines = append(routines, routine)
	}
	maxPages := (total + res - 1) / res
	return routines, maxPages, nil
}

func GetUserRoutines(adminId string) (routines []models.Routine, err error) {
	rows, err := db.Query("SELECT * FROM routines WHERE admin_id = $1", adminId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var routine models.Routine
		if err := rows.Scan(&routine.Id, &routine.AdminId, &routine.Name, &routine.Description, &routine.Img); err != nil {
			return nil, err
		}
		routines = append(routines, routine)
	}
	return routines, nil
}
func ILIKE(lhs, rhs StringExpression) BoolExpression {
	return BoolExp(CustomExpression(lhs, Token("ILIKE"), rhs))
}
func GetRoutinesJet(query string, page int64) (routines []model.Routines, maxPages int64, err error) {
	var res int64 = 20
	pageOffset := res * page
	queryPattern := "%" + query + "%"
	stmt := SELECT(
		Routines.AllColumns,
		Raw("COUNT(*) OVER() AS total_count"),
	).FROM(
		Routines,
	).WHERE(
		ILIKE(Routines.Name, String(queryPattern)),
	).LIMIT(res).OFFSET(pageOffset)
	type RoutineWithCount struct {
		model.Routines
		TotalCount int64 `db:"total_count"`
	}
	var dest []RoutineWithCount
	err = stmt.Query(db, &dest)
	if err != nil {
		return nil, 0, err
	}
	if len(dest) > 0 {
		total := dest[0].TotalCount
		maxPages = (total + res - 1) / res
	}
	routines = make([]model.Routines, len(dest))
	for i, row := range dest {
		routines[i] = row.Routines
	}
	return routines, maxPages, nil
}
