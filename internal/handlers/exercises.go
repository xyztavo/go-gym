package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/xyztavo/go-gym/internal/database"
	"github.com/xyztavo/go-gym/internal/models"
	"github.com/xyztavo/go-gym/internal/utils"
)

func CreateExercise(w http.ResponseWriter, r *http.Request) {
	newExercise := new(models.CreateExercise)
	if err := utils.BindAndValidate(r, newExercise); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newExerciseId, err := database.CreateExercise(newExercise)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	m := map[string]string{
		"message":       "exercise created with ease!",
		"newExerciseId": newExerciseId,
	}
	b, _ := json.Marshal(m)
	w.WriteHeader(http.StatusCreated)
	w.Write(b)
}

func GetExercises(w http.ResponseWriter, r *http.Request) {
	exercises, err := database.GetExercises()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, _ := json.Marshal(exercises)
	w.Write(b)
}
