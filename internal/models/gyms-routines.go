package models

type CreateGymRoutine struct {
	RoutineId string `json:"routineId"  validate:"required"`
}

type GymRoutine struct {
	Id        string `json:"id"`
	RoutineId string `json:"routineId"`
	GymId     string `json:"gymId"`
}
