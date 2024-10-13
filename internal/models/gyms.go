package models

type Gym struct {
	AdminId     string `json:"adminId"`
	Id          string `json:"id"`
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
