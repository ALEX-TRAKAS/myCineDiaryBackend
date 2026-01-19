package services

import (
	"mycinediarybackend/models"
	"mycinediarybackend/repositories"
	"time"
)

func AddThreadPost(threadPost *models.ThreadPost) error {
	return repositories.AddThreadPost(threadPost)
}

func RemoveThreadPost(threadPostID uint64) error {
	return repositories.RemoveThreadPost(threadPostID)
}

func GetThreadPostsByThreadID(threadID uint64) ([]models.ThreadPost, error) {
	return repositories.GetThreadPostsByThreadID(threadID)
}

func UpdateThreadPostBody(threadPostID uint64, newBody string, editedAt time.Time) error {
	return repositories.UpdateThreadPostBody(threadPostID, newBody, editedAt)
}
