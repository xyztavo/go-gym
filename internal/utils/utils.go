package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

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
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Token expires in 24 hours
			},
		})

	tokenString, err := token.SignedString([]byte(configs.GetJWTSecret()))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func UserIdFromToken(r *http.Request) (userId string, err error) {
	header := r.Header.Get("Authorization")
	authType := strings.Split(header, " ")[0]
	if header == "" {
		return "", errors.New("no auth header set")
	}
	if authType != "Bearer" {
		return "", errors.New("auth type is not bearer")
	}
	tokenString := strings.Split(header, " ")[1]
	claims := new(models.UserJwt)
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(configs.GetJWTSecret()), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}))

	if err != nil {
		return "", fmt.Errorf("invalid token: %w", err)
	}
	if !token.Valid {
		return "", errors.New("token is not valid")
	}
	return claims.Id, nil
}
