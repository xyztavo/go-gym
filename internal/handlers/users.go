package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

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
