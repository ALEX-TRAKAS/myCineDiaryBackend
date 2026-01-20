package services

import (
	"context"
	"errors"
	"mycinediarybackend/models"
	"mycinediarybackend/repositories"
	"mycinediarybackend/utils"
)

func Register(ctx context.Context, username, email, password string) (*models.User, error) {
	hashed, _ := utils.HashPassword(password)

	user := &models.User{
		Username:     username,
		Email:        email,
		PasswordHash: hashed,
	}

	if err := repositories.CreateUser(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}

func Login(ctx context.Context, email, password string) (*models.User, error) {
	user, err := repositories.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if !utils.CheckPassword(password, user.PasswordHash) {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func GetUserByID(ctx context.Context, id uint64) (*models.User, error) {
	return repositories.GetUserByID(ctx, id)
}
