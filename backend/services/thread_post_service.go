package services

import (
	"context"
	"mycinediarybackend/models"
	"mycinediarybackend/repositories"
	"time"
)

func AddThreadPost(ctx context.Context, threadPost *models.ThreadPost) error {
	return repositories.AddThreadPost(ctx, threadPost)
}

func RemoveThreadPost(ctx context.Context, threadPostID uint64) error {
	return repositories.RemoveThreadPost(ctx, threadPostID)
}

func GetThreadPostsByThreadID(ctx context.Context, threadID uint64) ([]models.ThreadPost, error) {
	return repositories.GetThreadPostsByThreadID(ctx, threadID)
}

func UpdateThreadPostBody(ctx context.Context, threadPostID uint64, newBody string, editedAt time.Time) error {
	return repositories.UpdateThreadPostBody(ctx, threadPostID, newBody, editedAt)
}
