package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/xyztavo/go-gym/internal/database"
	"github.com/xyztavo/go-gym/internal/models"
	"github.com/xyztavo/go-gym/internal/utils"
)

func CreateGym(w http.ResponseWriter, r *http.Request) {
	createGymBody := new(models.CreateGym)
	if err := utils.BindAndValidate(r, createGymBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userIdFromToken := utils.UserIdFromToken(r)
	createdGymId, status, err := database.CreateGym(userIdFromToken, createGymBody)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}
	m := map[string]string{
		"message":      "gym created with ease",
		"createdGymId": createdGymId,
	}
	b, _ := json.Marshal(m)
	w.WriteHeader(status)
	w.Write(b)
}
