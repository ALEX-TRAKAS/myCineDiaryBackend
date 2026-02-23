package services

import (
	"context"
	"mycinediarybackend/models"
	"mycinediarybackend/repositories"
)

func AddUserMedia(ctx context.Context, userMedia *models.UserMedia) error {
	return repositories.AddUserMedia(ctx, userMedia)
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
