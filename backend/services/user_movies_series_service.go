package services

import (
	"mycinediarybackend/models"
	"mycinediarybackend/repositories"
)

func AddUserMovie(userMovie *models.UserMovie) error {
	return repositories.AddUserMovie(userMovie)
}

func RemoveUserMovie(userID uint64, tmdbMovieID int) error {
	return repositories.RemoveUserMovie(userID, tmdbMovieID)
}

func AddUserSeries(userSeries *models.UserSeries) error {
	return repositories.AddUserSeries(userSeries)
}

func RemoveUserSeries(userID uint64, tmdbSeriesID int) error {
	return repositories.RemoveUserSeries(userID, tmdbSeriesID)
}

func GetUserMovies(userID uint64) ([]models.UserMovie, error) {
	return repositories.GetUserMovies(userID)
}

func GetUserSeries(userID uint64) ([]models.UserSeries, error) {
	return repositories.GetUserSeries(userID)
}
