package models

import "time"

type Plan struct {
	Id          string  `json:"id"`
	GymId       string  `json:"gymId"`
	AdminId     string  `json:"adminId"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Duration    int     `json:"duration"`
	Img         string  `json:"img"`
}

type CreatePlan struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float32 `json:"price" validate:"required"`
	Duration    int     `json:"duration" validate:"required"`
	Img         string  `json:"img" validate:"required"`
}

type SetUserPlan struct {
	Email  string `json:"email" validate:"required"`
	PlanId string `json:"planId" validate:"required"`
}
type UpdatePlan struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float32 `json:"price" validate:"required"`
	Duration    int     `json:"duration" validate:"required"`
	Img         string  `json:"img" validate:"required"`
}
type UserPlanDetails struct {
	Name        *string    `json:"name"`
	Description *string    `json:"description"`
	Duration    *int       `json:"duration"`
	Img         *string    `json:"img"`
	LastPayment *time.Time `json:"lastPayment"`
	ExpiresIn   *float64   `json:"expiresIn"`
	ExpiresAt   *time.Time `json:"expiresAt"`
}
