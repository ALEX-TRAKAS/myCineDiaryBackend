package services

import (
	"mycinediarybackend/models"
	"mycinediarybackend/repositories"
)

func CreateUser(user *models.User) error {
	return repositories.CreateUser(user)
}

func GetUser(email string) (*models.User, error) {
	return repositories.GetUserByEmail(email)
}
