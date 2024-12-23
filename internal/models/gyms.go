package models

type Gym struct {
	Id          string `json:"id"`
	AdminId     string `json:"adminId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Number      string `json:"number"`
	Img         string `json:"img"`
}
type CreateGym struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Location    string `json:"location" validate:"required"`
	Number      string `json:"number" validate:"required"`
	Img         string `json:"img" validate:"required"`
}
type GymDetails struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Location    string        `json:"location"`
	Number      string        `json:"number"`
	Image       string        `json:"image"`
	Plans       []GymPlans    `json:"plans"`
	Routines    []GymRoutines `json:"routines"`
}

type GymPlans struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Duration    int     `json:"duration"`
	Img         string  `json:"img"`
}

type GymRoutines struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Img         string `json:"img"`
}

type UpdateGym struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Location    string `json:"location" validate:"required"`
	Number      string `json:"number" validate:"required"`
	Img         string `json:"img" validate:"required"`
}
