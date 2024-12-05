package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/xyztavo/go-gym/internal/database"
	"github.com/xyztavo/go-gym/internal/utils"
)

func CreateGymRoutine(w http.ResponseWriter, r *http.Request) {
	id := utils.UserIdFromToken(r)
	routineId := chi.URLParam(r, "id")
	user, err := database.GetUserById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	createdGymRoutineId, err := database.CreateGymRoutine(*user.GymId, routineId)
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
