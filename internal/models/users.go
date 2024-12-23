package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserJwt struct {
	Id string `json:"id"`
	jwt.RegisteredClaims
}

type User struct {
	Id          string     `json:"id"`
	GymId       *string    `json:"gymId"`
	Name        string     `json:"name"`
	Email       string     `json:"email"`
	Role        string     `json:"role"`
	Password    string     `json:"password"`
	PlanId      *string    `json:"planId"`
	LastPayment *time.Time `json:"lastPayment"`
	CreatedAt   time.Time  `json:"createdAt"`
}

type SetUserGymAdmin struct {
	Id string `json:"id" validate:"required"`
}
type SetUserGymAdminByEmail struct {
	Email string `json:"email" validate:"required"`
}
type SetGymUser struct {
	Id string `json:"id" validate:"required"`
}

type SetGymUserByEmail struct {
	Email string `json:"email" validate:"required"`
}

type CreateUser struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required|email"`
	Password string `json:"password" validate:"required"`
}

type AuthUser struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
