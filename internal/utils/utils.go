package utils

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/xyztavo/go-gym/internal/configs"
	"github.com/xyztavo/go-gym/internal/models"
)

func BindAndValidate(r *http.Request, structs any) error {
	if err := json.NewDecoder(r.Body).Decode(&structs); err != nil {
		return err
	}
	validate := validator.New()
	if err := validate.Struct(structs); err != nil {
		return err
	}
	return nil
}

func CreateUserJwt(id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		models.UserJwt{
			Id: id,
		})

	tokenString, err := token.SignedString([]byte(configs.GetJWTSecret()))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
