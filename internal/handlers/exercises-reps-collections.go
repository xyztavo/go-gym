package handlers

import (
	"encoding/json"
	"net/http"

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
	createdRoutineCollectionRoutineId, err := database.CreateExercisesRepCollection(exerciseRepsCollectionBody)
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

func GetExercisesRepsCollections(w http.ResponseWriter, r *http.Request) {
	routinesCollectionRoutines, err := database.GetExercisesRoutinesCollectionsRoutines()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, _ := json.Marshal(routinesCollectionRoutines)
	w.Write(b)
}
