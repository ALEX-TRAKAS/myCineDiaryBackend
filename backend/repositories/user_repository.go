package repositories

import (
	"database/sql"
	"errors"

	"mycinediarybackend/database"
	"mycinediarybackend/models"
)

func CreateUser(user *models.User) error {
	query := `
		INSERT INTO users (username, email, password_hash)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	err := database.DB.
		QueryRow(query, user.Username, user.Email, user.PasswordHash).
		Scan(&user.ID)

	return err
}

func GetUserByEmail(email string) (*models.User, error) {
	query := `
		SELECT id, username, email, password_hash
		FROM users
		WHERE email = $1
	`

	var user models.User
	err := database.DB.
		QueryRow(query, email).
		Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash)

	if err == sql.ErrNoRows {
		return nil, errors.New("invalid credentials")
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}
