package database

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/xyztavo/go-gym/internal/models"
)

func CreatePlan(gymId string, plan *models.CreatePlan) (createdPlanId string, err error) {
	id, _ := gonanoid.New()
	err = db.QueryRow("INSERT INTO plans (id, gym_id, name, description, price, duration) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", id, gymId, plan.Name, plan.Description, plan.Price, plan.Duration).Scan(&createdPlanId)
	if err != nil {
		return "", err
	}
	return createdPlanId, nil
}
