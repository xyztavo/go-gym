package utils

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gookit/validate"
	"github.com/xyztavo/go-gym/internal/configs"
	"github.com/xyztavo/go-gym/internal/models"
)

func BindAndValidate(r *http.Request, structs any) error {
	if err := json.NewDecoder(r.Body).Decode(&structs); err != nil {
		return err
	}
	v := validate.Struct(structs)
	if !v.Validate() {
		return v.Errors
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

func UserIdFromToken(r *http.Request) (userId string) {
	header := r.Header.Get("Authorization")
	authType := strings.Split(header, " ")[0]
	if header == "" {
		return ""
	}
	if authType != "Bearer" {
		return ""
	}
	tokenString := strings.Split(header, " ")[1]
	claims := new(models.UserJwt)
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(configs.GetJWTSecret()), nil
	})
	if err != nil {
		return ""
	}
	if !token.Valid {
		return ""
	}
	return claims.Id
}
