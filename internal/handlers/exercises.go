package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

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
	query := r.URL.Query().Get("query")
	page := r.URL.Query().Get("page")
	intPage, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, err.Error()+" should specify page", http.StatusBadRequest)
		return
	}
	exercises, err := database.GetExercises(query, intPage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, _ := json.Marshal(exercises)
	w.Write(b)
}
