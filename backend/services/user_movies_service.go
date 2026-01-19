package services

import (
	"context"
	"mycinediarybackend/models"
	"mycinediarybackend/repositories"
)

func AddUserMovie(ctx context.Context, userMovie *models.UserMovie) error {
	return repositories.AddUserMovie(ctx, userMovie)
}

func RemoveUserMovie(ctx context.Context, userID uint64, tmdbMovieID int) error {
	return repositories.RemoveUserMovie(ctx, userID, tmdbMovieID)
}

func GetUserMovies(ctx context.Context, userID uint64) ([]models.UserMovie, error) {
	return repositories.GetUserMovies(ctx, userID)
}
