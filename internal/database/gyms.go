package database

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/xyztavo/go-gym/internal/models"
)

func CreateGym(userId string, createGym *models.CreateGym) (createdGymId string, statusCode int, err error) {
	user, err := GetUserById(userId)
	if err != nil {
		return "", http.StatusNotFound, errors.New("user does not exists")
	}
	if user.GymId != nil {
		return "", http.StatusBadRequest, errors.New("user already is in a gym")
	}
	id, _ := gonanoid.New()
	err = db.QueryRow("INSERT INTO gyms (admin_id, id, name, description, location, number, img) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		user.Id, id, createGym.Name, createGym.Description, createGym.Location, createGym.Number, createGym.Img).Scan(&createdGymId)
	if err != nil {
		return "", http.StatusInternalServerError, fmt.Errorf("could not create reason: %v", err.Error())
	}
	_, err = db.Exec("UPDATE users SET gym_id = $1 WHERE id = $2", createdGymId, user.Id)
	if err != nil {
		return "", http.StatusInternalServerError, errors.New("could not update admin user reason")
	}
	return createdGymId, http.StatusCreated, nil
}

func GetGymById(gymId string) (gym models.Gym, err error) {
	err = db.QueryRow("SELECT * FROM gyms WHERE id = $1", gymId).Scan(&gym.Id, &gym.AdminId, &gym.Name, &gym.Description, &gym.Location, &gym.Number, &gym.Img)
	if err != nil {
		return gym, err
	}
	return gym, nil
}

func SetGymUser(userId string, adminId string) (status int, err error) {
	_, err = GetUserById(userId)
	if err != nil {
		return http.StatusNotFound, errors.New("cannot find user with id :" + userId)
	}
	gym, err := GetUserGym(adminId)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	_, err = db.Exec("UPDATE users SET gym_id = $1 WHERE id = $2", gym.Id, userId)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func SetGymUserByEmail(email string, adminId string) (status int, err error) {
	_, err = GetUserByEmail(email)
	if err != nil {
		return http.StatusNotFound, errors.New("cannot find user with email :" + email)
	}
	gym, err := GetUserGym(adminId)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	_, err = db.Exec("UPDATE users SET gym_id = $1 WHERE email = $2", gym.Id, email)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func GetUserGymDetails(userId string) (gymDetails models.GymDetails, err error) {
	rows, err := db.Query(`
	SELECT g.name AS gym_name, g.description AS gym_description, 
		g.location AS gym_location, g.number AS gym_number, g.img AS gym_image, p.id AS plan_id, p.name AS plan_name, p.description AS plan_description, 
		p.price AS plan_price, p.duration AS plan_duration, p.img AS plan_image, r.id AS gym_routine_id, r.name AS gym_routine_name, r.description AS gym_routine_description,
		r.img AS routine_img
		FROM users AS u 
		LEFT JOIN gyms AS g ON u.gym_id = g.id 
		LEFT JOIN plans AS p ON p.gym_id = g.id 
		LEFT JOIN gyms_routines AS gr ON gr.gym_id = g.id
		LEFT JOIN routines AS r ON gr.routine_id = r.id
		WHERE u.id = $1
	`, userId)
	if err != nil {
		return gymDetails, err
	}
	plansMap := make(map[string]models.GymPlans)
	routinesMap := make(map[string]models.GymRoutines)

	for rows.Next() {
		var (
			gymName, gymDescription, gymLocation, gymNumber, gymImage string
			planId, planName, planDescription, planImg                sql.NullString
			planPrice                                                 sql.NullFloat64
			planDuration                                              sql.NullInt64
			routineId, routineName, routineDescription, routineImg    sql.NullString
		)

		if err := rows.Scan(&gymName, &gymDescription, &gymLocation, &gymNumber, &gymImage, &planId,
			&planName, &planDescription, &planPrice, &planDuration, &planImg,
			&routineId, &routineName, &routineDescription, &routineImg); err != nil {
			return gymDetails, err
		}
		// Set gym-level details only once
		if gymDetails.Name == "" {
			gymDetails.Name = gymName
			gymDetails.Description = gymDescription
			gymDetails.Location = gymLocation
			gymDetails.Number = gymNumber
			gymDetails.Image = gymImage
		}

		// Handle plans
		if planName.Valid {
			if _, exists := plansMap[planName.String]; !exists {
				plan := models.GymPlans{
					Id:          planId.String,
					Name:        planName.String,
					Description: planDescription.String,
					Price:       planPrice.Float64,
					Duration:    int(planDuration.Int64),
					Img:         planImg.String,
				}
				plansMap[planName.String] = plan
				gymDetails.Plans = append(gymDetails.Plans, plan)
			}
		}

		// Handle routines
		if routineName.Valid {
			if _, exists := routinesMap[routineName.String]; !exists {
				routine := models.GymRoutines{
					Id:          routineId.String,
					Name:        routineName.String,
					Description: routineDescription.String,
					Img:         routineImg.String,
				}
				routinesMap[routineName.String] = routine
				gymDetails.Routines = append(gymDetails.Routines, routine)
			}
		}
	}

	// Check for errors after row iteration
	if err := rows.Err(); err != nil {
		return gymDetails, err
	}

	return gymDetails, nil
}

func UpdateGym(adminId string, gym *models.UpdateGym) (status int, err error) {
	_, err = GetUserGym(adminId)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	_, err = db.Exec("UPDATE gyms SET name = $1, description = $2, location = $3, number = $4, img = $5 WHERE admin_id = $6", gym.Name, gym.Description, gym.Location, gym.Number, gym.Img, adminId)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}
