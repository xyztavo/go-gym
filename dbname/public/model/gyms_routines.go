//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type GymsRoutines struct {
	ID        string `sql:"primary_key"`
	RoutineID string
	GymID     string
}