package database

import (
	"errors"
	"fmt"
	"math"
	"time"

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
	err = db.QueryRow("SELECT gyms.id, gyms.admin_id, gyms.name, gyms.description, gyms.location, gyms.number, gyms.img  FROM gyms JOIN users ON gyms.id = users.gym_id WHERE users.id = $1",
		userId).Scan(&gym.Id, &gym.AdminId, &gym.Name, &gym.Description, &gym.Location, &gym.Number, &gym.Img)
	if err != nil {
		return gym, err
	}
	return gym, nil
}

func SetUserPlan(setUserPlan *models.SetUserPlan) (err error) {
	user, err := GetUserById(setUserPlan.UserId)
	if err != nil {
		return err
	}
	// if user already has paid
	if user.LastPayment != nil {
		var plan models.Plan
		err = db.QueryRow("SELECT * FROM plans WHERE id = $1", *user.PlanId).Scan(&plan.Id, &plan.GymId, &plan.Name, &plan.Description, &plan.Price, &plan.Duration, &plan.Img)
		if err != nil {
			return err
		}
		// get date until expiration
		dateUntilExpires := user.LastPayment.AddDate(0, 0, plan.Duration+1)
		timeComparison := dateUntilExpires.Compare(time.Now())
		// if plan has expired, set the current timestamp to be the last payment
		if timeComparison < 0 {
			_, err = db.Exec("UPDATE users SET plan_id = $1, last_payment = current_timestamp WHERE id = $2", setUserPlan.PlanId, setUserPlan.UserId)
			if err != nil {
				return err
			}
			return nil
		}
		// if plan has not expired yet, set last payment the exact date that it will expire
		_, err = db.Exec("UPDATE users SET plan_id = $1, last_payment = $2 WHERE id = $3", setUserPlan.PlanId, dateUntilExpires, setUserPlan.UserId)
		if err != nil {
			return err
		}
		return nil
	} else {
		_, err = db.Exec("UPDATE users SET plan_id = $1, last_payment = current_timestamp WHERE id = $2", setUserPlan.PlanId, setUserPlan.UserId)
		if err != nil {
			return err
		}
		return nil
	}
}

func GetGymUsers(gymId string) (users []models.User, err error) {
	rows, err := db.Query("SELECT * FROM users WHERE gym_id = $1", gymId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var user models.User
		rows.Scan(&user.Id, &user.GymId, &user.Name, &user.Email, &user.Role, &user.Password, &user.PlanId, &user.LastPayment, &user.CreatedAt)
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func CheckIn(userId string) (daysUntilPlanExpires float64, err error) {
	user, err := GetUserById(userId)
	if err != nil {
		return 0, errors.New("could not find user")
	}
	var plan models.Plan
	err = db.QueryRow("SELECT * FROM plans WHERE id = $1", *user.PlanId).Scan(&plan.Id, &plan.GymId, &plan.Name, &plan.Description, &plan.Price, &plan.Duration, &plan.Img)
	if err != nil {
		return 0, err
	}
	dateUntilExpires := user.LastPayment.AddDate(0, 0, plan.Duration+1)
	fmt.Println(dateUntilExpires)
	timeComparison := dateUntilExpires.Compare(time.Now())
	daysUntilExpiration := math.Floor(time.Until(dateUntilExpires).Hours() / 24)
	if timeComparison < 0 {
		return 0, fmt.Errorf("your plan has expired in %v days, please renew it", daysUntilExpiration)
	}
	return daysUntilExpiration, nil
}

func GetUserPlanDetails(userId string) (userPlanDetails models.UserPlanDetails, err error) {
	if err = db.QueryRow(`
	SELECT p.name, p.description, p.duration, p.img, u.last_payment FROM users AS u 
	LEFT JOIN plans AS p ON p.id = u.plan_id WHERE u.id = $1
	`, userId).Scan(&userPlanDetails.Name, &userPlanDetails.Description, &userPlanDetails.Duration, &userPlanDetails.Img, &userPlanDetails.LastPayment); err != nil {
		return userPlanDetails, err
	}
	if userPlanDetails.Name == nil {
		return userPlanDetails, errors.New("user does not have plan")
	}
	dateUntilExpires := userPlanDetails.LastPayment.AddDate(0, 0, *userPlanDetails.Duration+1)
	userPlanDetails.ExpiresAt = &dateUntilExpires
	daysUntilExpiration := math.Floor(time.Until(dateUntilExpires).Hours() / 24)
	userPlanDetails.ExpiresIn = &daysUntilExpiration
	return userPlanDetails, nil
}
