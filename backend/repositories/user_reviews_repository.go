package repositories

import (
	"context"
	"mycinediarybackend/database"
	"mycinediarybackend/models"
)

func AddReview(ctx context.Context, review *models.Review) error {
	query := `
		INSERT INTO reviews (user_id, tmdb_id, media_type, rating, review_text, is_spoiler)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (user_id, tmdb_id, media_type) 
		DO UPDATE SET rating = EXCLUDED.rating, review_text = EXCLUDED.review_text, is_spoiler = EXCLUDED.is_spoiler, updated_at = NOW()
	`
	_, err := database.DB.Exec(ctx, query,
		review.UserID,
		review.TMDBID,
		review.MediaType,
		review.Rating,
		review.ReviewText,
		review.IsSpoiler,
	)
	return err
}

func RemoveReview(ctx context.Context, userID uint64, tmdbID int, mediaType string) error {
	query := `
		DELETE FROM reviews
		WHERE user_id = $1 AND tmdb_id = $2 AND media_type = $3
	`
	_, err := database.DB.Exec(ctx, query, userID, tmdbID, mediaType)
	return err
}

func GetReviews(ctx context.Context, userID uint64, page int, limit int, mediaType string) (*models.PaginatedReviews, error) {
	if page < 1 {
		page = 1
	}
	if limit <= 0 {
		limit = 12
	}
	offset := (page - 1) * limit

	var totalItems int
	countQuery := `SELECT COUNT(*) FROM reviews WHERE user_id = $1`
	args := []interface{}{userID}
	if mediaType != "" {
		countQuery += " AND media_type = $2"
		args = append(args, mediaType)
	}
	if err := database.DB.QueryRow(ctx, countQuery, args...).Scan(&totalItems); err != nil {
		return nil, err
	}

	query := `SELECT id, user_id, tmdb_id, media_type, rating, review_text, is_spoiler, created_at, updated_at
	          FROM reviews
			  WHERE user_id = $1`
	args = []interface{}{userID}
	if mediaType != "" {
		query += " AND media_type = $2"
		args = append(args, mediaType)
	}
	query += " ORDER BY updated_at DESC LIMIT $3 OFFSET $4"
	args = append(args, limit, offset)

	rows, err := database.DB.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []models.Review
	for rows.Next() {
		var r models.Review
		if err := rows.Scan(&r.ID, &r.UserID, &r.TMDBID, &r.MediaType, &r.Rating, &r.ReviewText, &r.IsSpoiler, &r.CreatedAt, &r.UpdatedAt); err != nil {
			return nil, err
		}
		reviews = append(reviews, r)
	}

	totalPages := (totalItems + limit - 1) / limit
	return &models.PaginatedReviews{
		Reviews:     reviews,
		CurrentPage: page,
		TotalPages:  totalPages,
		TotalItems:  totalItems,
	}, nil
}
