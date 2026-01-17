package services

import (
	"errors"

	"mycinediarybackend/models"
	"mycinediarybackend/repositories"
	"mycinediarybackend/utils"
)

func Register(username, email, password string) error {
	hash, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user := &models.User{
		Username:     username,
		Email:        email,
		PasswordHash: hash,
	}

	return repositories.CreateUser(user)
}

func Login(email, password string) (*models.User, error) {
	user, err := repositories.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if !utils.CheckPassword(password, user.PasswordHash) {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}
