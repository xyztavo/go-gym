package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/xyztavo/go-gym/internal/database"
	"github.com/xyztavo/go-gym/internal/models"
	"github.com/xyztavo/go-gym/internal/utils"
)

func CreateRoutineExerciseRepsCollection(w http.ResponseWriter, r *http.Request) {
	createRoutineExerciseRepsCollectionBody := new(models.CreateRoutineExerciseRepsCollection)
	if err := utils.BindAndValidate(r, createRoutineExerciseRepsCollectionBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdRoutineExerciseRepsCollectionId, err := database.CreateRoutineExerciseRepsCollection(createRoutineExerciseRepsCollectionBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	m := map[string]string{
		"message":                                "created routine exercise reps collection with ease!",
		"createdRoutineExerciseRepsCollectionId": createdRoutineExerciseRepsCollectionId,
	}
	b, _ := json.Marshal(m)
	w.WriteHeader(http.StatusCreated)
	w.Write(b)
}
