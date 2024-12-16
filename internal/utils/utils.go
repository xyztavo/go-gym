package utils

import (
	"encoding/json"
	"net/http"
	"strings"

	. "github.com/go-jet/jet/v2/postgres"
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

func ILIKE(lhs, rhs StringExpression) BoolExpression {
	return BoolExp(CustomExpression(lhs, Token("ILIKE"), rhs))
}
