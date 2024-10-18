package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/xyztavo/go-gym/internal/database"
	"github.com/xyztavo/go-gym/internal/models"
	"github.com/xyztavo/go-gym/internal/utils"
)

func CreatePlan(w http.ResponseWriter, r *http.Request) {
	idFromToken := utils.UserIdFromToken(r)
	planBody := new(models.CreatePlan)
	if err := utils.BindAndValidate(r, planBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	gym, err := database.GetUserGym(idFromToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	createdPlanId, err := database.CreatePlan(gym.Id, planBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	m := map[string]string{
		"message":       "created plan with ease",
		"gymId":         gym.Id,
		"createdPlanId": createdPlanId,
	}
	b, _ := json.Marshal(m)
	w.WriteHeader(http.StatusCreated)
	w.Write(b)
}

func GetUserGymPlans(w http.ResponseWriter, r *http.Request) {
	idFromToken := utils.UserIdFromToken(r)
	gym, err := database.GetUserGym(idFromToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	plans, err := database.GetGymPlans(gym.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, _ := json.Marshal(plans)
	w.Write(b)
}
