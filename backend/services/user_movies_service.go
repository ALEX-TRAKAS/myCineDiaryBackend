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

func GetUserMovies(ctx context.Context, userID uint64, page int, limit int) ([]models.UserMovie, error) {
	paginatedUserMovies, err := repositories.GetUserMovies(ctx, userID, page, limit)
	if err != nil {
		return nil, err
	}
	return paginatedUserMovies.Movies, nil
}
