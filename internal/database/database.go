package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/xyztavo/go-gym/internal/configs"
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
	// Migrate tables
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS gyms (
	    id VARCHAR(40) PRIMARY KEY,
	    admin_id VARCHAR(40) UNIQUE NOT NULL,
	    name VARCHAR(40) UNIQUE NOT NULL,
	    description VARCHAR(200) NOT NULL,
	    location VARCHAR(40) NOT NULL,
	    number VARCHAR(40) NOT NULL,
	    img VARCHAR(200) NOT NULL
	);
	CREATE TABLE IF NOT EXISTS plans (
	    id VARCHAR(40) PRIMARY KEY,
	    gym_id VARCHAR(40) NOT NULL,
	    admin_id VARCHAR(40) NOT NULL,
	    name VARCHAR(40) UNIQUE NOT NULL,
	    description VARCHAR(200) NOT NULL,
	    price DOUBLE PRECISION NOT NULL,
	    duration INT NOT NULL,
	    img VARCHAR(200) NOT NULL,
	    FOREIGN KEY (gym_id) REFERENCES gyms(id) ON DELETE CASCADE
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
		FOREIGN KEY (gym_id) REFERENCES gyms(id) ON DELETE SET NULL,
		FOREIGN KEY (plan_id) REFERENCES plans(id) ON DELETE SET NULL
	);
	ALTER TABLE gyms ADD FOREIGN KEY (admin_id) REFERENCES users(id) ON DELETE SET NULL;
	ALTER TABLE plans ADD FOREIGN KEY (admin_id) REFERENCES users(id) ON DELETE CASCADE;
	CREATE TABLE IF NOT EXISTS exercises (
		id VARCHAR(40) PRIMARY KEY,
		name VARCHAR(100) UNIQUE NOT NULL,
		description VARCHAR(200) NOT NULL,
		gif VARCHAR(200) NOT NULL
	);
	CREATE TABLE IF NOT EXISTS collections (
		id VARCHAR(40) PRIMARY KEY,
		admin_id VARCHAR(40) NOT NULL,
		name VARCHAR(40) UNIQUE NOT NULL,
		description VARCHAR(200) NOT NULL,
		img VARCHAR(200) NOT NULL,
		FOREIGN KEY (admin_id) REFERENCES users(id) ON DELETE CASCADE
	);
	CREATE TABLE IF NOT EXISTS exercises_reps_collections (
		id VARCHAR(40) PRIMARY KEY,
		admin_id VARCHAR(40) NOT NULL,
		collection_id VARCHAR(40) NOT NULL,
		exercise_id VARCHAR(40) NOT NULL,
		reps INT NOT NULL,
		sets INT NOT NULL,
		FOREIGN KEY (collection_id) REFERENCES collections(id) ON DELETE CASCADE,
		FOREIGN KEY (exercise_id) REFERENCES exercises(id) ON DELETE CASCADE,
		FOREIGN KEY (admin_id) REFERENCES users(id) ON DELETE CASCADE
	);
	CREATE TABLE IF NOT EXISTS routines (
		id VARCHAR(40) PRIMARY KEY,
		admin_id VARCHAR(40) NOT NULL,
		name VARCHAR(40) UNIQUE NOT NULL,
		description VARCHAR(200) NOT NULL,
		img VARCHAR(200) NOT NULL,
		FOREIGN KEY (admin_id) REFERENCES users(id) ON DELETE CASCADE
	);
	CREATE TABLE IF NOT EXISTS routines_collections (
		id VARCHAR(40) PRIMARY KEY,
		admin_id VARCHAR(40) NOT NULL,
		routine_id VARCHAR(40) NOT NULL,
		collection_id VARCHAR(40) NOT NULL,
		FOREIGN KEY (routine_id) REFERENCES routines(id) ON DELETE CASCADE,
		FOREIGN KEY (collection_id) REFERENCES collections(id) ON DELETE CASCADE,
		FOREIGN KEY (admin_id) REFERENCES users(id) ON DELETE CASCADE
	);
	CREATE TABLE IF NOT EXISTS gyms_routines (
		id VARCHAR(40) PRIMARY KEY,
		routine_id VARCHAR(40) NOT NULL,
		gym_id VARCHAR(40) NOT NULL,
		FOREIGN KEY (routine_id) REFERENCES routines(id) ON DELETE CASCADE,
		FOREIGN KEY (gym_id) REFERENCES gyms(id) ON DELETE CASCADE
	);
	`)
	if err != nil {
		return err
	}
	return nil
}
