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
	// Migrate tables
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS gyms (
	id VARCHAR(40) PRIMARY KEY,
	admin_id VARCHAR(40) UNIQUE NOT NULL,
	name VARCHAR(40) UNIQUE NOT NULL,
	description VARCHAR(40) NOT NULL,
	location VARCHAR(40) NOT NULL,
	number VARCHAR(40) NOT NULL
	);
	CREATE TABLE IF NOT EXISTS plans (
	id VARCHAR(40) PRIMARY KEY,
	gym_id VARCHAR(40) NOT NULL,
	name VARCHAR(40) UNIQUE NOT NULL,
	description VARCHAR(40) NOT NULL,
	price DOUBLE PRECISION NOT NULL,
	duration INT NOT NULL,
	FOREIGN KEY (gym_id) REFERENCES gyms(id)
	);
	CREATE TABLE IF NOT EXISTS users (
	id VARCHAR(40) PRIMARY KEY,
	gym_id VARCHAR(40),
	name VARCHAR(40) NOT NULL,
	email VARCHAR(40) UNIQUE NOT NULL, 
	role VARCHAR(40) NOT NULL,
	password VARCHAR(200) NOT NULL,
	plan_id VARCHAR(40),
	last_payment TIMESTAMP,
	created_at TIMESTAMP DEFAULT NOW(),
	FOREIGN KEY (gym_id) REFERENCES gyms(id),
	FOREIGN KEY (plan_id) REFERENCES plans(id)
	);
	ALTER TABLE gyms ADD FOREIGN KEY (admin_id) REFERENCES users(id);
	CREATE TABLE IF NOT EXISTS exercises (
	id VARCHAR(40) PRIMARY KEY,
	name VARCHAR(40) UNIQUE NOT NULL,
	description VARCHAR(40) NOT NULL,
	gif VARCHAR(40) NOT NULL
	);
	CREATE TABLE IF NOT EXISTS routines (
	id VARCHAR(40) PRIMARY KEY,
	name VARCHAR(40) UNIQUE NOT NULL,
	description VARCHAR(40) NOT NULL,
	thumb VARCHAR(40) NOT NULL
	);
	CREATE TABLE IF NOT EXISTS exercises_routines (
    id VARCHAR(40) PRIMARY KEY,
    routine_id VARCHAR(40) NOT NULL,
    exercise_id VARCHAR(40) UNIQUE NOT NULL,
    reps INT NOT NULL,
	sets INT NOT NULL,
    FOREIGN KEY (routine_id) REFERENCES routines(id),
    FOREIGN KEY (exercise_id) REFERENCES exercises(id)
	);
	`)
	if err != nil {
		return err
	}
	// create admin user
	db.Exec(`
	INSERT INTO users (id, name, email, role, password) VALUES ($1, $2, $3, 'admin', $4);
	`, id, user.Name, user.Email, HashedPassword)
	return nil
}
