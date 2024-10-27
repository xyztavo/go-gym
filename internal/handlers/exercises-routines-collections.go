package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/xyztavo/go-gym/internal/database"
	"github.com/xyztavo/go-gym/internal/models"
	"github.com/xyztavo/go-gym/internal/utils"
)

func CreateExercisesRoutineCollection(w http.ResponseWriter, r *http.Request) {
	createRoutineColletionBody := new(models.CreateExerciseRoutineColletion)
	if err := utils.BindAndValidate(r, createRoutineColletionBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdRoutineCollectionId, err := database.CreateExercisesRoutineCollection(createRoutineColletionBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	m := map[string]string{
		"message":                   "routine collection created with ease",
		"createdRoutineColletionId": createdRoutineCollectionId,
	}
	b, _ := json.Marshal(m)
	w.WriteHeader(http.StatusCreated)
	w.Write(b)
}

func GetExercisesRoutinesCollections(w http.ResponseWriter, r *http.Request) {
	routinesCollections, err := database.GetExercisesRoutinesCollections()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, _ := json.Marshal(routinesCollections)
	w.Write(b)
}

func CreateExercisesRoutineCollectionRoutine(w http.ResponseWriter, r *http.Request) {
	routineCollectionRoutineBody := new(models.CreateExerciseRoutineCollectionRoutine)
	if err := utils.BindAndValidate(r, routineCollectionRoutineBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdRoutineCollectionRoutineId, err := database.CreateExercisesRoutineCollectionRoutine(*routineCollectionRoutineBody)
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

func GetExercisesRoutinesCollectionsRoutines(w http.ResponseWriter, r *http.Request) {
	routinesCollectionRoutines, err := database.GetExercisesRoutinesCollectionsRoutines()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, _ := json.Marshal(routinesCollectionRoutines)
	w.Write(b)
}
