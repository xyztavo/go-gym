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

func CreateExercise(w http.ResponseWriter, r *http.Request) {
	idFromToken := utils.UserIdFromToken(r)
	newExercise := new(models.CreateExercise)
	if err := utils.BindAndValidate(r, newExercise); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newExerciseId, err := database.CreateExercise(idFromToken, newExercise)
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
	exercises, maxPages, err := database.GetExercises(query, intPage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := map[string]interface{}{
		"exercises": exercises,
		"maxPages":  maxPages,
		"page":      intPage,
	}
	b, _ := json.Marshal(response)
	w.Write(b)
}

func GetUserExercises(w http.ResponseWriter, r *http.Request) {
	idFromToken := utils.UserIdFromToken(r)
	exercises, err := database.GetUserExercises(idFromToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, _ := json.Marshal(exercises)
	w.Write(b)
}

func GetExerciseById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	exercise, err := database.GetExerciseById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	b, _ := json.Marshal(exercise)
	w.Write(b)
}

func UpdateExercise(w http.ResponseWriter, r *http.Request) {
	idFromToken := utils.UserIdFromToken(r)
	id := chi.URLParam(r, "id")
	newExercise := new(models.UpdateExercise)
	if err := utils.BindAndValidate(r, newExercise); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := database.UpdateExercise(idFromToken, id, newExercise)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	m := map[string]string{
		"message": "exercise updated with ease!",
	}
	b, _ := json.Marshal(m)
	w.Write(b)
}

func DeleteExercise(w http.ResponseWriter, r *http.Request) {
	idFromToken := utils.UserIdFromToken(r)
	id := chi.URLParam(r, "id")
	err := database.DeleteExercise(idFromToken, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	m := map[string]string{
		"message": "exercise deleted with ease!",
	}
	b, _ := json.Marshal(m)
	w.Write(b)
}
