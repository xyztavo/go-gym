package database

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/xyztavo/go-gym/internal/models"
)

func GetPlanById(planId string) (plan models.Plan, err error) {
	err = db.QueryRow("SELECT * FROM plans WHERE id = $1", planId).Scan(&plan.Id, &plan.GymId, &plan.Name, &plan.Description, &plan.Price, &plan.Duration, &plan.Img)
	if err != nil {
		return models.Plan{}, err
	}
	return plan, nil
}
func CreatePlan(gymId string, plan *models.CreatePlan) (createdPlanId string, err error) {
	id, _ := gonanoid.New()
	err = db.QueryRow("INSERT INTO plans (id, gym_id, name, description, price, duration, img) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		id, gymId, plan.Name, plan.Description, plan.Price, plan.Duration, plan.Img).Scan(&createdPlanId)
	if err != nil {
		return "", err
	}
	return createdPlanId, nil
}

func GetGymPlans(gymId string) (plans []models.Plan, err error) {
	rows, err := db.Query("SELECT * FROM plans WHERE gym_id = $1", gymId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var plan models.Plan
		rows.Scan(&plan.Id, &plan.GymId, &plan.Name, &plan.Description, &plan.Price, &plan.Duration, &plan.Img)
		plans = append(plans, plan)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return plans, nil
}

func DeleteGymPlan(planId string) (err error) {
	_, err = db.Exec("DELETE FROM plans WHERE id = $1", planId)
	if err != nil {
		return err
	}
	return nil
}
