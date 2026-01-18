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

func GetUserMovies(userID uint64) ([]models.UserMovie, error) {
	return repositories.GetUserMovies(userID)
}
