package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/xyztavo/go-gym/internal/database"
	"github.com/xyztavo/go-gym/internal/utils"
)

func CreateGymRoutine(w http.ResponseWriter, r *http.Request) {
	idFromToken, err := utils.UserIdFromToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	routineId := chi.URLParam(r, "id")
	user, err := database.GetUserById(idFromToken)
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
	idFromToken, err := utils.UserIdFromToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	user, err := database.GetUserById(idFromToken)
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

func DeleteGymRoutine(w http.ResponseWriter, r *http.Request) {
	idFromToken, err := utils.UserIdFromToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	id := chi.URLParam(r, "id")
	err = database.DeleteGymRoutine(idFromToken, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
