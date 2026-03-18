package services

import (
	"context"
	"log"
	"mycinediarybackend/models"
	"mycinediarybackend/repositories"
)

func AddUserMedia(ctx context.Context, userMedia *models.UserMedia, genres []models.Genre) error {
	if err := repositories.AddUserMedia(ctx, userMedia, genres); err != nil {
		return err
	}

	tmdbID := userMedia.TMDBID

	if err := CreateActivity(
		ctx,
		userMedia.UserID,
		models.ActivityAddWatchlist,
		&tmdbID,
		userMedia.Rating,
	); err != nil {
		log.Println("Failed to create activity:", err)

	}
	return nil
}

func RemoveUserMedia(ctx context.Context, userID uint64, tmdbID int, mediaType models.MediaType) error {
	return repositories.RemoveUserMedia(ctx, userID, tmdbID, string(mediaType))
}

func GetUserMedia(ctx context.Context, userID uint64, page, limit int, mediaType models.MediaType) (*models.PaginatedUserMedia, error) {
	paginatedUserMedia, err := repositories.GetUserMedia(ctx, userID, string(mediaType), page, limit)
	if err != nil {
		return nil, err
	}
	return paginatedUserMedia, nil
}

func GetUserMediaByTMDBID(ctx context.Context, userID uint64, tmdbID int, mediaType models.MediaType) (*models.UserMedia, error) {
	userMedia, err := repositories.GetUserMediaByTMDBID(ctx, userID, tmdbID, string(mediaType))
	if err != nil {
		return nil, err
	}
	return userMedia, nil
}
