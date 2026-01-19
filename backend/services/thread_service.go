package services

import (
	"mycinediarybackend/models"
	"mycinediarybackend/repositories"
)

func GetAllThreads() ([]models.Thread, error) {
	return repositories.GetAllThreads()
}

func CreateThread(thread models.Thread) error {
	return repositories.CreateThread(&thread)
}
func GetThreadByID(id string) (*models.Thread, error) {
	return repositories.GetThreadByID(id)
}

func DeleteThread(id string) error {
	return repositories.DeleteThread(id)
}
