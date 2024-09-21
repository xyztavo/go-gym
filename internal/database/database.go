package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/xyztavo/go-gym/internal/configs"
	"github.com/xyztavo/go-gym/internal/models"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", configs.GetDBConnStr())
	if err != nil {
		log.Fatal(err)
	}
}

func Migrate() error {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
	id VARCHAR(40) PRIMARY KEY,
	name VARCHAR(40) NOT NULL,
	email VARCHAR(40) UNIQUE NOT NULL, 
	password VARCHAR(200) NOT NULL
	);
	`)
	if err != nil {
		return err
	}
	return nil
}

func CreateUser(user *models.CreateUser) (id string, err error) {
	id, _ = gonanoid.New()
	HashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	_, err = db.Exec(`INSERT INTO users (id, name, email, password) VALUES ($1, $2, $3, $4)`, id, &user.Name, &user.Email, HashedPassword)
	if err != nil {
		return "", err
	}
	return id, nil
}

func GetUserByEmail(email string) (user models.User, err error) {
	if err := db.QueryRow("SELECT * FROM users WHERE email = $1", email).Scan(&user.Id, &user.Name, &user.Email, &user.Password); err != nil {
		return user, err
	}
	return user, nil
}
