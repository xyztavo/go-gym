package models

type CreateGymRoutine struct {
	Name        string `json:"name" validate:"required" `
	Description string `json:"description"  validate:"required"`
	RoutineId   string `json:"routineId"  validate:"required"`
}

type GymRoutine struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	RoutineId   string `json:"routineId"`
	GymId       string `json:"gymId"`
}
