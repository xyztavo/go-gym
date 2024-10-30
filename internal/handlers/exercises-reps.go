package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/xyztavo/go-gym/internal/database"
	"github.com/xyztavo/go-gym/internal/models"
	"github.com/xyztavo/go-gym/internal/utils"
)

func CreateExerciseReps(w http.ResponseWriter, r *http.Request) {
	exerciseRepsBody := new(models.CreateExerciseReps)
	if err := utils.BindAndValidate(r, exerciseRepsBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdExerciseRepsId, err := database.CreateExerciseReps(exerciseRepsBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	m := map[string]string{
		"message":                  "created exercises reps with ease",
		"createdExerciseRoutineId": createdExerciseRepsId,
	}
	b, _ := json.Marshal(m)
	w.WriteHeader(http.StatusCreated)
	w.Write(b)
}

func GetExercisesRepsByCollectionId(w http.ResponseWriter, r *http.Request) {
	getExercisesBody := new(models.GetExercisesRepsByCollectionId)
	if err := utils.BindAndValidate(r, getExercisesBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	exercises, err := database.GetExercisesRepsByCollectionId(getExercisesBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, _ := json.Marshal(exercises)
	w.Write(b)
}
