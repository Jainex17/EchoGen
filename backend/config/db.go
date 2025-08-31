package config

import (
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDatabase() {
	connStr := DatabaseURL

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	DB = db.DB
	log.Println("Database connected successfully")

	// err = runMigrations()
	// if err != nil {
	// 	log.Fatal("error running migrations: %w", err)
	// }
	// log.Println("Migrations run successfully")
}

// func runMigrations() error {
// 	// Create users table
// 	_, err := DB.Exec(`
// 	CREATE TABLE IF NOT EXISTS users (
// 		id SERIAL PRIMARY KEY,
// 		email VARCHAR(255) UNIQUE NOT NULL,
// 		name VARCHAR(255) NOT NULL,
// 		picture VARCHAR(500),
// 		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
// 	);
// 	`)
// 	if err != nil {
// 		return err
// 	}

// 	_, err = DB.Exec(`
// 		CREATE TABLE IF NOT EXISTS audio_generations (
// 			id SERIAL PRIMARY KEY,
// 			user_id INTEGER REFERENCES users(id),
// 			prompt TEXT NOT NULL,
// 			content TEXT NOT NULL,
// 			audio_url VARCHAR(500),
// 			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
// 		);
// 	`)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func CloseDatabase() {
	if DB != nil {
		DB.Close()
	}
}
