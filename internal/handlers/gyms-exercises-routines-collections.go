package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/xyztavo/go-gym/internal/database"
	"github.com/xyztavo/go-gym/internal/models"
	"github.com/xyztavo/go-gym/internal/utils"
)

func CreateGymExerciseCollection(w http.ResponseWriter, r *http.Request) {
	userId := utils.UserIdFromToken(r)
	user, err := database.GetUserById(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	createGymExerciseCollectionBody := new(models.CreateGymExerciseRoutineCollection)
	if err := utils.BindAndValidate(r, createGymExerciseCollectionBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdGymExerciseCollectionId, err := database.CreateGymExerciseRoutineCollection(*user.GymId, createGymExerciseCollectionBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	m := map[string]string{
		"message":                       "created gym exercise collection with ease",
		"createdGymRoutineCollectionId": createdGymExerciseCollectionId,
	}
	b, _ := json.Marshal(m)
	w.WriteHeader(http.StatusCreated)
	w.Write(b)
}
