package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/xyztavo/go-gym/internal/database"
	"github.com/xyztavo/go-gym/internal/models"
	"github.com/xyztavo/go-gym/internal/utils"
)

func CreateCollection(w http.ResponseWriter, r *http.Request) {
	idFromToken := utils.UserIdFromToken(r)
	createRoutineColletionBody := new(models.CreateCollection)
	if err := utils.BindAndValidate(r, createRoutineColletionBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdRoutineCollectionId, err := database.CreateCollection(idFromToken, createRoutineColletionBody)
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

func GetCollections(w http.ResponseWriter, r *http.Request) {
	routinesCollections, err := database.GetCollections()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, _ := json.Marshal(routinesCollections)
	w.Write(b)
}

func GetAdminCollections(w http.ResponseWriter, r *http.Request) {
	id := utils.UserIdFromToken(r)
	routinesCollections, err := database.GetAdminCollections(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, _ := json.Marshal(routinesCollections)
	w.Write(b)
}

func GetCollectionsByRoutineId(w http.ResponseWriter, r *http.Request) {
	getCollectionsByRoutineIdBody := new(models.GetCollectionsByRoutineId)
	if err := utils.BindAndValidate(r, getCollectionsByRoutineIdBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	collections, err := database.GetCollectionsByRoutineId(getCollectionsByRoutineIdBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	b, _ := json.Marshal(collections)
	w.Write(b)
}
