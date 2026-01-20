package repositories

import (
	"context"
	"database/sql"
	"errors"

	"mycinediarybackend/database"
	"mycinediarybackend/models"
)

func CreateUser(ctx context.Context, user *models.User) error {
	query := `
        INSERT INTO users (username, email, password_hash)
        VALUES ($1, $2, $3)
        RETURNING id, created_at
    `
	return database.DB.QueryRow(ctx, query, user.Username, user.Email, user.PasswordHash).
		Scan(&user.ID, &user.CreatedAt)

}

func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `
        SELECT id, username, email, password_hash, created_at
        FROM users
        WHERE email = $1
    `

	var user models.User
	err := database.DB.QueryRow(ctx, query, email).
		Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, errors.New("invalid credentials")
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserByID(ctx context.Context, id uint64) (*models.User, error) {
	query := `
        SELECT id, username, email, created_at
        FROM users
        WHERE id = $1
    `

	var user models.User
	err := database.DB.QueryRow(ctx, query, id).
		Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}
