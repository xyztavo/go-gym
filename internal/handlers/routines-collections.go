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
	idFromToken, err := utils.UserIdFromToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
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

func DeleteRoutineFromCollection(w http.ResponseWriter, r *http.Request) {
	routineId := chi.URLParam(r, "routineId")
	routineCollectionId := chi.URLParam(r, "routineCollectionId")
	idFromToken, err := utils.UserIdFromToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	err = database.DeleteRoutineFromCollection(idFromToken, routineId, routineCollectionId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	m := map[string]string{
		"message": "deleted routine from collection with ease!",
	}
	b, _ := json.Marshal(m)
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
