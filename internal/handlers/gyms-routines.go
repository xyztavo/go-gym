package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/xyztavo/go-gym/internal/database"
	"github.com/xyztavo/go-gym/internal/models"
	"github.com/xyztavo/go-gym/internal/utils"
)

func CreateGymRoutine(w http.ResponseWriter, r *http.Request) {
	id := utils.UserIdFromToken(r)
	createdGymRoutineBody := new(models.CreateGymRoutine)
	if err := utils.BindAndValidate(r, createdGymRoutineBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := database.GetUserById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	createdGymRoutineId, err := database.CreateGymRoutine(*user.GymId, createdGymRoutineBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	m := map[string]string{
		"message":             "created gym routine with ease",
		"createdGymRoutineId": createdGymRoutineId,
	}
	b, _ := json.Marshal(m)
	w.WriteHeader(http.StatusCreated)
	w.Write(b)
}

func GetUserGymRoutines(w http.ResponseWriter, r *http.Request) {
	id := utils.UserIdFromToken(r)
	user, err := database.GetUserById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if user.GymId == nil {
		http.Error(w, "user is not in a gym", http.StatusBadRequest)
		return
	}
	gymRoutines, err := database.GetGymRoutines(*user.GymId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	b, _ := json.Marshal(gymRoutines)
	w.Write(b)
}
