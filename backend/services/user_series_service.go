package services

import (
	"context"
	"mycinediarybackend/models"
	"mycinediarybackend/repositories"
)

func AddUserSeries(ctx context.Context, userSeries *models.UserSeries) error {
	return repositories.AddUserSeries(ctx, userSeries)
}

func RemoveUserSeries(ctx context.Context, userID uint64, tmdbSeriesID int) error {
	return repositories.RemoveUserSeries(ctx, userID, tmdbSeriesID)
}

func GetUserSeries(ctx context.Context, userID uint64, page int, limit int) ([]models.UserSeries, error) {
	paginatedUserSeries, err := repositories.GetUserSeries(ctx, userID, page, limit)
	if err != nil {
		return nil, err
	}
	return paginatedUserSeries.Series, nil
}
