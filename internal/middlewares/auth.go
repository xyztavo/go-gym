package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/xyztavo/go-gym/internal/configs"
	"github.com/xyztavo/go-gym/internal/database"
	"github.com/xyztavo/go-gym/internal/models"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		authType := strings.Split(header, " ")[0]
		if header == "" {
			http.Error(w, "no auth header set", http.StatusUnauthorized)
			return
		}
		if authType != "Bearer" {
			http.Error(w, "auth type is not bearer", http.StatusUnauthorized)
			return
		}
		tokenString := strings.Split(header, " ")[1]
		claims := new(models.UserJwt)
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(configs.GetJWTSecret()), nil
		})
		if err != nil {
			http.Error(w, "token is not valid", http.StatusUnauthorized)
			return
		}
		if !token.Valid {
			http.Error(w, "token is not valid", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func AdminAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		authType := strings.Split(header, " ")[0]
		if header == "" {
			http.Error(w, "no auth header set", http.StatusUnauthorized)
			return
		}
		if authType != "Bearer" {
			http.Error(w, "auth type is not bearer", http.StatusUnauthorized)
			return
		}
		tokenString := strings.Split(header, " ")[1]
		claims := new(models.UserJwt)
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(configs.GetJWTSecret()), nil
		})
		if err != nil {
			http.Error(w, "token is not valid", http.StatusUnauthorized)
			return
		}
		if !token.Valid {
			http.Error(w, "token is not valid", http.StatusUnauthorized)
			return
		}
		user, err := database.GetUserById(claims.Id)
		if err != nil {
			http.Error(w, fmt.Sprintf("admin user with id %v not found, reason: %v", user.Id, err.Error()), http.StatusUnauthorized)
			return
		}
		if user.Role != "admin" {
			http.Error(w, "user not admin", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
