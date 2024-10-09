package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/xyztavo/go-gym/internal/configs"
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
	user := configs.GetAdminInfo()
	id, _ := gonanoid.New()
	HashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
	id VARCHAR(40) PRIMARY KEY,
	name VARCHAR(40) NOT NULL,
	email VARCHAR(40) UNIQUE NOT NULL, 
	role VARCHAR(40) NOT NULL,
	password VARCHAR(200) NOT NULL
	);
	CREATE TABLE IF NOT EXISTS exercises (
	id VARCHAR(40) PRIMARY KEY,
	name VARCHAR(40) NOT NULL,
	description VARCHAR(40) NOT NULL,
	gif VARCHAR(40) NOT NULL
	);
	CREATE TABLE IF NOT EXISTS routines (
	id VARCHAR(40) PRIMARY KEY,
	name VARCHAR(40) NOT NULL,
	description VARCHAR(40) NOT NULL,
	thumb VARCHAR(40) NOT NULL
	);
	CREATE TABLE IF NOT EXISTS exercises_routines (
    id VARCHAR(40) PRIMARY KEY,
    routine_id VARCHAR(40) NOT NULL,
    exercise_id VARCHAR(40) NOT NULL,
    reps INT NOT NULL,
    FOREIGN KEY (routine_id) REFERENCES routines(id),
    FOREIGN KEY (exercise_id) REFERENCES exercises(id)
	);
	`)
	if err != nil {
		return err
	}
	// create admin user
	_, err = db.Exec(`
	INSERT INTO users (id, name, email, role, password) VALUES ($1, $2, $3, 'admin', $4);
	`, id, user.Name, user.Email, HashedPassword)
	if err != nil {
		return nil
	}
	return nil
}
