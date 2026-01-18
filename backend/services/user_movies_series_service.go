package services

import (
	"mycinediarybackend/models"
	"mycinediarybackend/repositories"
)

func AddUserMovie(userMovie *models.UserMovie) error {
	return repositories.AddUserMovie(userMovie)
}
func AddUserSeries(userSeries *models.UserSeries) error {
	return repositories.AddUserSeries(userSeries)
}

func GetUserMovies(userID uint64) ([]models.UserMovie, error) {
	return repositories.GetUserMovies(userID)
}

func GetUserSeries(userID uint64) ([]models.UserSeries, error) {
	return repositories.GetUserSeries(userID)
}
