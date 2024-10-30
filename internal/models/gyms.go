package models

type Gym struct {
	Id          string `json:"id"`
	AdminId     string `json:"adminId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Number      string `json:"number"`
}
type CreateGym struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Location    string `json:"location" validate:"required"`
	Number      string `json:"number" validate:"required"`
}
type GymDetails struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Location    string        `json:"location"`
	Plans       []GymPlans    `json:"plans"`
	Routines    []GymRoutines `json:"routines"`
}

type GymPlans struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Duration    int     `json:"duration"`
}

type GymRoutines struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
