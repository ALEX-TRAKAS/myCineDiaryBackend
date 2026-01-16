package repositories

import (
	"mycinediarybackend/database"
	"mycinediarybackend/models"
)

func CreateUser(user *models.User) error {
	query := `
		INSERT INTO users (name, email)
		VALUES ($1, $2)
		RETURNING id
	`

	return database.DB.
		QueryRow(query, user.Username, user.Email).
		Scan(&user.ID)
}

func GetUserByID(id int64) (*models.User, error) {
	query := `
		SELECT id, name, email
		FROM users
		WHERE id = $1
	`

	var user models.User
	err := database.DB.
		QueryRow(query, id).
		Scan(&user.ID, &user.Username, &user.Email)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
