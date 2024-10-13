package database

import (
	"errors"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/xyztavo/go-gym/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user *models.CreateUser) (id string, err error) {
	id, _ = gonanoid.New()
	HashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	_, err = db.Exec(`INSERT INTO users (id, name, email, role, password) VALUES ($1, $2, $3, $4, $5)`, id, &user.Name, &user.Email, "regular", HashedPassword)
	if err != nil {
		return "", err
	}
	return id, nil
}

func GetUserByEmail(email string) (user models.User, err error) {
	if err := db.QueryRow("SELECT * FROM users WHERE email = $1", email).Scan(&user.Id, &user.GymId, &user.Name, &user.Email, &user.Role, &user.Password, &user.PlanId, &user.LastPayment, &user.CreatedAt); err != nil {
		return user, err
	}
	return user, nil
}

func GetUserById(id string) (user models.User, err error) {
	if err := db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&user.Id, &user.GymId, &user.Name, &user.Email, &user.Role, &user.Password, &user.PlanId, &user.LastPayment, &user.CreatedAt); err != nil {
		return user, err
	}
	return user, nil
}

func SetUserGymAdmin(id string) error {
	_, err := GetUserById(id)
	if err != nil {
		return errors.New("user not found")
	}
	_, err = db.Exec("UPDATE users SET role = 'gym-admin' WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func GetUserGym(userId string) (gym models.Gym, err error) {
	user, err := GetUserById(userId)
	if err != nil {
		return gym, err
	}
	gym, err = GetGymById(*user.GymId)
	if err != nil {
		return gym, err
	}
	return gym, nil
}
