package handlers

import (
	"encoding/json"
	"net/http"

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
	createdRoutineBodyId, err := database.CreateRoutine(createRoutineBody)
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
	routines, err := database.GetRoutines()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, _ := json.Marshal(routines)
	w.Write(b)
}
