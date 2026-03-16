package services

import (
	"context"
	"log"
	"mycinediarybackend/models"
	"mycinediarybackend/repositories"
)

func AddReview(ctx context.Context, review *models.Review) error {
	if err := repositories.AddReviewAndUpdateMedia(ctx, review); err != nil {
		return err
	}

	tmdbID := review.TMDBID
	rating := review.Rating

	if err := CreateActivity(
		ctx,
		review.UserID,
		models.ActivityReviewMovie,
		&tmdbID,
		&rating,
	); err != nil {
		log.Println("Failed to create activity:", err)

	}

	return nil
}

func RemoveReview(ctx context.Context, userID uint64, tmdbID int, mediaType string) error {
	return repositories.RemoveReview(ctx, userID, tmdbID, mediaType)
}

func GetReviews(ctx context.Context, userID uint64, page int, limit int, mediaType string) ([]models.Review, error) {
	paginated, err := repositories.GetReviews(ctx, userID, page, limit, mediaType)
	if err != nil {
		return nil, err
	}
	return paginated.Reviews, nil
}

func GetMediaReviewsPublic(ctx context.Context, tmdbID int, page int, limit int, mediaType string) ([]models.Review, error) {
	paginated, err := repositories.GetMediaReviewsPublic(ctx, tmdbID, mediaType, page, limit)
	if err != nil {
		return nil, err
	}
	return paginated.Reviews, nil
}
