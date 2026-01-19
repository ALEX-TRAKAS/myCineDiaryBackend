package services

import (
	"context"
	"mycinediarybackend/models"
	"mycinediarybackend/repositories"
)

func GetAllThreads(ctx context.Context) ([]models.Thread, error) {
	return repositories.GetAllThreads(ctx)
}

func CreateThread(ctx context.Context, thread models.Thread) error {
	return repositories.CreateThread(ctx, &thread)
}
func GetThreadByID(ctx context.Context, id string) (*models.Thread, error) {
	return repositories.GetThreadByID(ctx, id)
}

func DeleteThread(ctx context.Context, id string) error {
	return repositories.DeleteThread(ctx, id)
}
