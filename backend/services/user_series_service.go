package services

import (
	"mycinediarybackend/models"
	"mycinediarybackend/repositories"
)

func AddUserSeries(userSeries *models.UserSeries) error {
	return repositories.AddUserSeries(userSeries)
}

func RemoveUserSeries(userID uint64, tmdbSeriesID int) error {
	return repositories.RemoveUserSeries(userID, tmdbSeriesID)
}

func GetUserSeries(userID uint64) ([]models.UserSeries, error) {
	return repositories.GetUserSeries(userID)
}
