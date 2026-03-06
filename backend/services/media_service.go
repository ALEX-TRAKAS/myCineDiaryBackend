package services

import (
	"context"
	"mycinediarybackend/models"
	"mycinediarybackend/repositories"
)

func GetMediaByTMDBID(ctx context.Context, tmdbID int, mediaType models.MediaType) (*models.Media, error) {
	media, err := repositories.GetMediaByTMDBID(ctx, tmdbID, string(mediaType))
	if err != nil {
		return nil, err
	}
	return media, nil
}
