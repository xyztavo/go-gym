package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/xyztavo/go-gym/internal/database"
	"github.com/xyztavo/go-gym/internal/models"
	"github.com/xyztavo/go-gym/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

func GetAuth(w http.ResponseWriter, r *http.Request) {
	authUserBody := new(models.AuthUser)
	if err := utils.BindAndValidate(r, authUserBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userFromDb, err := database.GetUserByEmail(authUserBody.Email)
	if err != nil {
		http.Error(w, "could not find user, err: "+err.Error(), http.StatusNotFound)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(userFromDb.Password), []byte(authUserBody.Password))
	if err != nil {
		http.Error(w, "password does not match", http.StatusUnauthorized)
		return
	}
	jwt, _ := utils.CreateUserJwt(userFromDb.Id)
	m := map[string]string{"message": "authorized", "token": jwt, "role": userFromDb.Role}
	b, _ := json.Marshal(m)
	w.Write(b)
}
