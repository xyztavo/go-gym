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
	query := r.URL.Query().Get("query")
	page := r.URL.Query().Get("page")
	intPage, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Invalid page parameter: "+err.Error(), http.StatusBadRequest)
		return
	}
	collections, maxPages, err := database.GetCollections(query, intPage)
	if err != nil {
		http.Error(w, "Error fetching collections: "+err.Error(), http.StatusInternalServerError)
		return
	}
	response := map[string]interface{}{
		"collections": collections,
		"maxPages":    maxPages,
		"page":        intPage,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
		return
	}
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
	id := chi.URLParam(r, "id")
	collections, err := database.GetCollectionsByRoutineId(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	b, _ := json.Marshal(collections)
	w.Write(b)
}
