package database

import (
	"errors"
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
	err = db.QueryRow("INSERT INTO gyms (admin_id, id, name, description, location, number) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		user.Id, id, createGym.Name, createGym.Description, createGym.Location, createGym.Number).Scan(&createdGymId)
	if err != nil {
		return "", http.StatusInternalServerError, errors.New("could not create ")
	}
	_, err = db.Exec("UPDATE users SET gym_id = $1 WHERE id = $2", createdGymId, user.Id)
	if err != nil {
		return "", http.StatusInternalServerError, errors.New("could not update admin user reason")
	}
	return createdGymId, http.StatusCreated, nil
}

func GetGymById(gymId string) (gym models.Gym, err error) {
	err = db.QueryRow("SELECT * FROM gyms WHERE id = $1", gymId).Scan(&gym.Id, &gym.AdminId, &gym.Name, &gym.Description, &gym.Location, &gym.Number)
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
