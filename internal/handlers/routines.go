package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

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
	idFromToken := utils.UserIdFromToken(r)
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
