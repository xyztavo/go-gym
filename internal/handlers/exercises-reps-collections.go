package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/xyztavo/go-gym/internal/database"
	"github.com/xyztavo/go-gym/internal/models"
	"github.com/xyztavo/go-gym/internal/utils"
)

func CreateExercisesRepsCollection(w http.ResponseWriter, r *http.Request) {
	exerciseRepsCollectionBody := new(models.CreateExerciseRepsCollection)
	if err := utils.BindAndValidate(r, exerciseRepsCollectionBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	idFromToken := utils.UserIdFromToken(r)
	createdRoutineCollectionRoutineId, err := database.CreateExercisesRepCollection(idFromToken, exerciseRepsCollectionBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	m := map[string]string{
		"message":                    "created routine collection routine with ease",
		"routineCollectionRoutineId": createdRoutineCollectionRoutineId,
	}
	b, _ := json.Marshal(m)
	w.WriteHeader(http.StatusCreated)
	w.Write(b)
}

func CreateMultipleExercisesRepCollection(w http.ResponseWriter, r *http.Request) {
	body := new(models.CreateMultipleExercisesRepCollection)
	if err := utils.BindAndValidate(r, body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	idFromToken := utils.UserIdFromToken(r)
	if err := database.CreateMultipleExercisesRepCollection(idFromToken, body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	m := map[string]string{"message": "exercises added to collection with ease"}
	b, _ := json.Marshal(m)
	w.WriteHeader(http.StatusCreated)
	w.Write(b)
}

func GetExercisesRepsCollections(w http.ResponseWriter, r *http.Request) {
	routinesCollectionRoutines, err := database.GetExercisesRepsCollection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, _ := json.Marshal(routinesCollectionRoutines)
	w.Write(b)
}

func GetExercisesRepsCollectionsByCollectionId(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	exercisesReps, err := database.GetExercisesRepsCollectionsByCollectionId(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, _ := json.Marshal(exercisesReps)
	w.Write(b)
}
