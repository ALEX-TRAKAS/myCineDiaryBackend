package services

import (
	"context"
	"time"

	"github.com/google/uuid"

	"mycinediarybackend/models"
	"mycinediarybackend/repositories"
)

func CreateActivity(
	ctx context.Context,
	userID uint64,
	activityType models.ActivityType,
	movieID *int,
	rating *int,
) error {

	activity := &models.Activity{
		ID:        uuid.New().String(),
		UserID:    userID,
		Type:      activityType,
		MovieID:   movieID,
		Rating:    rating,
		CreatedAt: time.Now(),
	}

	return repositories.CreateActivity(ctx, activity)
}

func GetUserActivity(ctx context.Context, userID uint64) ([]models.Activity, error) {
	return repositories.GetUserActivities(ctx, userID, 20)
}
