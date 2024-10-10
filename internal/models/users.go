package models

import "github.com/golang-jwt/jwt/v5"

type UserJwt struct {
	Id string `json:"id"`
	jwt.RegisteredClaims
}
type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Password string `json:"password"`
}

type SetUserGymAdmin struct {
	Id string `json:"id" validate:"required"`
}

type CreateUser struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthUser struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
