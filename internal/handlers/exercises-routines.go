package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/xyztavo/go-gym/internal/database"
	"github.com/xyztavo/go-gym/internal/models"
	"github.com/xyztavo/go-gym/internal/utils"
)

func CreateExerciseRoutine(w http.ResponseWriter, r *http.Request) {
	exerciseRoutineBody := new(models.CreateExerciseRoutine)
	if err := utils.BindAndValidate(r, exerciseRoutineBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdExerciseRoutineId, err := database.CreateExerciseRoutine(exerciseRoutineBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	m := map[string]string{
		"message":                  "exercise added to routine with ease",
		"createdExerciseRoutineId": createdExerciseRoutineId,
	}
	b, _ := json.Marshal(m)
	w.WriteHeader(http.StatusCreated)
	w.Write(b)
}
