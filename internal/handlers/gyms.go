package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/xyztavo/go-gym/internal/database"
	"github.com/xyztavo/go-gym/internal/models"
	"github.com/xyztavo/go-gym/internal/utils"
)

func CreateGym(w http.ResponseWriter, r *http.Request) {
	createGymBody := new(models.CreateGym)
	if err := utils.BindAndValidate(r, createGymBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userIdFromToken := utils.UserIdFromToken(r)
	createdGymId, status, err := database.CreateGym(userIdFromToken, createGymBody)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}
	m := map[string]string{
		"message":      "gym created with ease",
		"createdGymId": createdGymId,
	}
	b, _ := json.Marshal(m)
	w.WriteHeader(status)
	w.Write(b)
}

func SetGymUser(w http.ResponseWriter, r *http.Request) {
	idFromToken := utils.UserIdFromToken(r)
	setGymUser := new(models.SetGymUser)
	if err := utils.BindAndValidate(r, setGymUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	statusCode, err := database.SetGymUser(setGymUser.Id, idFromToken)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}
	m := map[string]string{"message": fmt.Sprintf("user with id: %v is now in your gym!", setGymUser.Id)}
	b, _ := json.Marshal(m)
	w.WriteHeader(statusCode)
	w.Write(b)
}

func SetGymUserByEmail(w http.ResponseWriter, r *http.Request) {
	idFromToken := utils.UserIdFromToken(r)
	setGymUser := new(models.SetGymUserByEmail)
	if err := utils.BindAndValidate(r, setGymUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	statusCode, err := database.SetGymUserByEmail(setGymUser.Email, idFromToken)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}
	m := map[string]string{"message": fmt.Sprintf("user with email: %v is now in your gym!", setGymUser.Email)}
	b, _ := json.Marshal(m)
	w.WriteHeader(statusCode)
	w.Write(b)
}

func GetUserGymDetails(w http.ResponseWriter, r *http.Request) {
	idFromToken := utils.UserIdFromToken(r)
	gymDetails, err := database.GetUserGymDetails(idFromToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, _ := json.Marshal(gymDetails)
	w.Write(b)
}

func UpdateGym(w http.ResponseWriter, r *http.Request) {
	idFromToken := utils.UserIdFromToken(r)
	updateGymBody := new(models.UpdateGym)
	if err := utils.BindAndValidate(r, updateGymBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	statusCode, err := database.UpdateGym(idFromToken, updateGymBody)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}
	m := map[string]string{"message": "gym updated with ease"}
	b, _ := json.Marshal(m)
	w.WriteHeader(statusCode)
	w.Write(b)
}
