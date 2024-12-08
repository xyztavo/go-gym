package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/xyztavo/go-gym/internal/database"
	"github.com/xyztavo/go-gym/internal/models"
	"github.com/xyztavo/go-gym/internal/utils"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	createUser := new(models.CreateUser)
	if err := utils.BindAndValidate(r, createUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newUserId, err := database.CreateUser(createUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jwt, _ := utils.CreateUserJwt(newUserId)
	m := map[string]string{"message": "user created with ease", "token": jwt}
	b, _ := json.Marshal(m)
	w.WriteHeader(http.StatusCreated)
	w.Write(b)
}
func GetUserGym(w http.ResponseWriter, r *http.Request) {
	idFromToken := utils.UserIdFromToken(r)
	gym, err := database.GetUserGym(idFromToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, _ := json.Marshal(gym)
	w.Write(b)
}

func SetUserGymAdmin(w http.ResponseWriter, r *http.Request) {
	setUserGymAdmin := new(models.SetUserGymAdmin)
	if err := utils.BindAndValidate(r, setUserGymAdmin); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := database.SetUserGymAdmin(setUserGymAdmin.Id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	m := map[string]string{"message": fmt.Sprintf("user with id %v is now gym admin", setUserGymAdmin.Id)}
	b, _ := json.Marshal(m)
	w.Write(b)
}

func SetUserGymAdminByEmail(w http.ResponseWriter, r *http.Request) {
	SetUserGymAdminByEmail := new(models.SetUserGymAdminByEmail)
	if err := utils.BindAndValidate(r, SetUserGymAdminByEmail); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := database.SetUserGymAdminByEmail(SetUserGymAdminByEmail.Email); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	m := map[string]string{"message": fmt.Sprintf("user with email %v is now gym admin", SetUserGymAdminByEmail.Email)}
	b, _ := json.Marshal(m)
	w.Write(b)
}

func GetGymUsers(w http.ResponseWriter, r *http.Request) {
	idFromToken := utils.UserIdFromToken(r)
	gym, err := database.GetUserGym(idFromToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	users, err := database.GetGymUsers(gym.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	b, _ := json.Marshal(users)
	w.Write(b)
}

func CheckIn(w http.ResponseWriter, r *http.Request) {
	idFromToken := utils.UserIdFromToken(r)
	daysUntilPlanExpires, err := database.CheckIn(idFromToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	m := map[string]any{
		"message":              "check in approved",
		"daysUntilPlanExpires": daysUntilPlanExpires,
	}
	b, _ := json.Marshal(m)
	w.Write(b)
}

func CheckInByUserId(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	daysUntilPlanExpires, err := database.CheckIn(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	m := map[string]any{
		"message":              "check in approved",
		"daysUntilPlanExpires": daysUntilPlanExpires,
	}
	b, _ := json.Marshal(m)
	w.Write(b)
}

func GetUserPlanDetails(w http.ResponseWriter, r *http.Request) {
	idFromToken := utils.UserIdFromToken(r)
	userPlanDetails, err := database.GetUserPlanDetails(idFromToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, _ := json.Marshal(userPlanDetails)
	w.Write(b)
}
