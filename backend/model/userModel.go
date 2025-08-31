package model

import (
	"time"

	"echogen/backend/config"
)

type User struct {
	ID        int       `db:"id"`
	Email     string    `db:"email"`
	Name      string    `db:"name"`
	Picture   string    `db:"picture"`
	CreatedAt time.Time `db:"created_at"`
}

type UserRepository struct{}

func (r *UserRepository) CreateUser(user *User) error {
	query := `
		INSERT INTO users (email, name, picture) VALUES ($1, $2, $3)
	`

	_, err := config.DB.Exec(query, user.Email, user.Name, user.Picture)
	return err
}

func (r *UserRepository) GetUserByEmail(email string) (*User, error) {
	var user User

	query := `
		SELECT * FROM users WHERE email = $1
	`

	err := config.DB.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Name, &user.Picture, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	if user.ID == 0 {
		return nil, nil
	}

	return &user, nil
}

func (r *UserRepository) UserExists(email string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`

	err := config.DB.QueryRow(query, email).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
