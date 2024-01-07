package db

import "time"

type User struct {
	ID        int
	Username  string
	Password  string
	Email     string
	CreatedAt time.Time
	UpdateAt  time.Time
}

func CreateUser(username, password, email string) (*User, error) {
	var user User
	err := db.QueryRow(`
		INSERT INTO users (username, password, email)
		VALUES ($1, $2, $3)
		RETURNING id, username, password, email, created_at, updated_at
	`, username, password, email).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.CreatedAt,
		&user.UpdateAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	err := db.QueryRow(`
		SELECT id, username, password, email, created_at, updated_at
		FROM users
		WHERE username = $1
	`, username).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.CreatedAt,
		&user.UpdateAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserById(id string) (*User, error) {
	var user User
	err := db.QueryRow(`
		SELECT id, username, password, email, created_at, updated_at
		FROM users
		WHERE id = $1
	`, id).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.CreatedAt,
		&user.UpdateAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
