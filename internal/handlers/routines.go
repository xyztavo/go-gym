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
	createdRoutineId, err := database.CreateRoutine(createRoutineBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	m := map[string]string{
		"message":          "routine created with ease",
		"createdRoutineId": createdRoutineId,
	}
	b, _ := json.Marshal(m)
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
