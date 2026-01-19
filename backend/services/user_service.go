package services

import (
	"context"
	"errors"

	"mycinediarybackend/models"
	"mycinediarybackend/repositories"
	"mycinediarybackend/utils"
)

func Register(ctx context.Context, username, email, password string) error {
	hash, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user := &models.User{
		Username:     username,
		Email:        email,
		PasswordHash: hash,
	}

	_, err = repositories.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func Login(ctx context.Context, email, password string) (*models.User, error) {
	user, err := repositories.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if !utils.CheckPassword(password, user.PasswordHash) {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func GetUserByID(ctx context.Context, id uint64) (*models.User, error) {
	return repositories.GetUserByID(ctx, id)
}
