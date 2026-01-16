package services

import (
	"mycinediarybackend/models"
	"mycinediarybackend/repositories"
)

func CreateUser(user *models.User) error {
	return repositories.CreateUser(user)
}

func GetUser(id int64) (*models.User, error) {
	return repositories.GetUserByID(id)
}
