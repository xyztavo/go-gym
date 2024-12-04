package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/xyztavo/go-gym/internal/database"
	"github.com/xyztavo/go-gym/internal/utils"
)

func CreateRoutineCollection(w http.ResponseWriter, r *http.Request) {
	routineId := chi.URLParam(r, "routineId")
	collectionId := chi.URLParam(r, "collectionId")
	idFromToken := utils.UserIdFromToken(r)
	createdRoutineExerciseRepsCollectionId, err := database.CreateRoutineCollection(idFromToken, routineId, collectionId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	m := map[string]string{
		"message":                    "created routine collection with ease!",
		"createdRoutineCollectionId": createdRoutineExerciseRepsCollectionId,
	}
	b, _ := json.Marshal(m)
	w.WriteHeader(http.StatusCreated)
	w.Write(b)
}
