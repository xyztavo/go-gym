package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/xyztavo/go-gym/internal/database"
	"github.com/xyztavo/go-gym/internal/models"
	"github.com/xyztavo/go-gym/internal/utils"
)

func CreateRoutine(w http.ResponseWriter, r *http.Request) {
	createRoutineBody := new(models.CreateRoutine)
	if err := utils.BindAndValidate(r, createRoutineBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	idFromToken, err := utils.UserIdFromToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	createdRoutineBodyId, err := database.CreateRoutine(idFromToken, createRoutineBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	m := map[string]string{
		"message":              "created routine with ease",
		"createdRoutineBodyId": createdRoutineBodyId,
	}
	b, _ := json.Marshal(m)
	w.WriteHeader(http.StatusCreated)
	w.Write(b)
}

func GetRoutines(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	page := r.URL.Query().Get("page")
	intPage, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Invalid page parameter: "+err.Error(), http.StatusBadRequest)
		return
	}
	routines, maxPages, err := database.GetRoutines(query, intPage)
	if err != nil {
		http.Error(w, "Error fetching routines: "+err.Error(), http.StatusInternalServerError)
		return
	}
	response := map[string]interface{}{
		"routines": routines,
		"maxPages": maxPages,
		"page":     intPage,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetUserRoutines(w http.ResponseWriter, r *http.Request) {
	idFromToken, err := utils.UserIdFromToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	routines, err := database.GetUserRoutines(idFromToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, _ := json.Marshal(routines)
	w.Write(b)
}

func GetRoutineById(w http.ResponseWriter, r *http.Request) {
	routineId := chi.URLParam(r, "id")
	routine, err := database.GetRoutineById(routineId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	b, _ := json.Marshal(routine)
	w.Write(b)
}
func UpdateRoutine(w http.ResponseWriter, r *http.Request) {
	idFromToken, err := utils.UserIdFromToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	routineId := chi.URLParam(r, "id")
	routineBody := new(models.UpdateRoutine)
	if err := utils.BindAndValidate(r, routineBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = database.UpdateRoutine(routineId, idFromToken, routineBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	m := map[string]string{"message": "updated routine with ease!"}
	b, _ := json.Marshal(m)
	w.Write(b)
}
func DeleteRoutine(w http.ResponseWriter, r *http.Request) {
	idFromToken, err := utils.UserIdFromToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	routineId := chi.URLParam(r, "id")
	err = database.DeleteRoutine(routineId, idFromToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	m := map[string]string{"message": "deleted routine with ease!"}
	b, _ := json.Marshal(m)
	w.Write(b)
}
