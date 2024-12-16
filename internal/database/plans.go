package database

import (
	"errors"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/xyztavo/go-gym/internal/models"
)

func GetPlanById(planId string) (plan models.Plan, err error) {
	err = db.QueryRow("SELECT * FROM plans WHERE id = $1", planId).Scan(&plan.Id, &plan.GymId, &plan.AdminId, &plan.Name, &plan.Description, &plan.Price, &plan.Duration, &plan.Img)
	if err != nil {
		return models.Plan{}, err
	}
	return plan, nil
}
func CreatePlan(gymId string, adminId string, plan *models.CreatePlan) (createdPlanId string, err error) {
	id, _ := gonanoid.New()
	err = db.QueryRow("INSERT INTO plans (id, gym_id, admin_id, name, description, price, duration, img) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id",
		id, gymId, adminId, plan.Name, plan.Description, plan.Price, plan.Duration, plan.Img).Scan(&createdPlanId)
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
		if err := rows.Scan(&plan.Id, &plan.GymId, &plan.AdminId, &plan.Name, &plan.Description, &plan.Price, &plan.Duration, &plan.Img); err != nil {
			return nil, err
		}
		plans = append(plans, plan)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return plans, nil
}

func UpdatePlan(planId string, adminId string, plan *models.UpdatePlan) (err error) {
	plans, err := GetPlanById(planId)
	if err != nil {
		return err
	}
	if plans.AdminId != adminId {
		return errors.New("you are not the admin of this plan")
	}
	_, err = db.Exec("UPDATE plans SET name = $1, description = $2, price = $3, duration = $4, img = $5 WHERE id = $6", plan.Name, plan.Description, plan.Price, plan.Duration, plan.Img, planId)
	if err != nil {
		return err
	}
	return nil
}

func DeleteGymPlan(planId string, adminId string) (err error) {
	plans, err := GetPlanById(planId)
	if err != nil {
		return err
	}
	if plans.AdminId != adminId {
		return errors.New("you are not the admin of this plan")
	}
	_, err = db.Exec("DELETE FROM plans WHERE id = $1", planId)
	if err != nil {
		return err
	}
	return nil
}
