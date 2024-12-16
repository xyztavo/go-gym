package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
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
	createdPlanId, err := database.CreatePlan(gym.Id, idFromToken, planBody)
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

func GetPlanById(w http.ResponseWriter, r *http.Request) {
	planId := chi.URLParam(r, "id")
	plan, err := database.GetPlanById(planId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	b, _ := json.Marshal(plan)
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

func SetUserPlan(w http.ResponseWriter, r *http.Request) {
	userPlan := new(models.SetUserPlan)
	if err := utils.BindAndValidate(r, userPlan); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := database.SetUserPlan(userPlan); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	m := map[string]string{"message": "user plan updated with ease!"}
	b, _ := json.Marshal(m)
	w.Write(b)
}

func UpdatePlan(w http.ResponseWriter, r *http.Request) {
	idFromToken := utils.UserIdFromToken(r)
	planId := chi.URLParam(r, "id")
	planBody := new(models.UpdatePlan)
	if err := utils.BindAndValidate(r, planBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := database.UpdatePlan(planId, idFromToken, planBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	m := map[string]string{"message": "updated plan with ease!"}
	b, _ := json.Marshal(m)
	w.Write(b)
}

func DeleteGymPlan(w http.ResponseWriter, r *http.Request) {
	idFromToken := utils.UserIdFromToken(r)
	planId := chi.URLParam(r, "id")
	err := database.DeleteGymPlan(planId, idFromToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	m := map[string]string{"message": "deleted plan with ease!"}
	b, _ := json.Marshal(m)
	w.Write(b)
}
