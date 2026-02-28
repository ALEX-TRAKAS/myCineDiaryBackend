package repositories

import (
	"context"
	"mycinediarybackend/database"
	"mycinediarybackend/models"

	"github.com/jackc/pgx/v5"
)

func AddReviewAndUpdateMedia(ctx context.Context, review *models.Review) error {

	tx, err := database.DB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	query := `
		INSERT INTO reviews (user_id, tmdb_id, media_type, rating, review_text, is_spoiler, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW())
		ON CONFLICT (user_id, tmdb_id, media_type)
		DO UPDATE SET rating = EXCLUDED.rating, review_text = EXCLUDED.review_text, is_spoiler = EXCLUDED.is_spoiler, updated_at = NOW()
	`
	_, err = tx.Exec(ctx, query,
		review.UserID,
		review.TMDBID,
		review.MediaType,
		review.Rating,
		review.ReviewText,
		review.IsSpoiler,
	)
	if err != nil {
		return err
	}

	aggQuery := `
		UPDATE media
		SET average_rating = sub.avg,
		    reviews_count = sub.count
		FROM (
			SELECT COALESCE(AVG(rating),0) AS avg,
			       COUNT(*) AS count
			FROM reviews
			WHERE tmdb_id = $1 AND media_type = $2
		) sub
		WHERE media.tmdb_id = $1 AND media.media_type = $2
	`
	_, err = tx.Exec(ctx, aggQuery, review.TMDBID, review.MediaType)
	if err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

func RemoveReview(ctx context.Context, userID uint64, tmdbID int, mediaType string) error {
	tx, err := database.DB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	delQuery := `
		DELETE FROM reviews
		WHERE user_id = $1 AND tmdb_id = $2 AND media_type = $3
	`
	_, err = tx.Exec(ctx, delQuery, userID, tmdbID, mediaType)
	if err != nil {
		return err
	}
	if err := tx.Commit(ctx); err != nil {
		return err
	}

	aggQuery := `
		UPDATE media
		SET average_rating = sub.avg,
		    reviews_count = sub.count
		FROM (
			SELECT COALESCE(AVG(rating),0) AS avg,
			       COUNT(*) AS count
			FROM reviews
			WHERE tmdb_id = $1 AND media_type = $2
		) sub
		WHERE media.tmdb_id = $1 AND media.media_type = $2
	`
	_, err = tx.Exec(ctx, aggQuery, tmdbID, mediaType)
	if err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

func GetReviews(
	ctx context.Context,
	userID uint64,
	page int,
	limit int,
	mediaType string,
) (*models.PaginatedReviews, error) {

	if page < 1 {
		page = 1
	}
	if limit <= 0 {
		limit = 12
	}
	offset := (page - 1) * limit

	countQuery := `
		SELECT COUNT(*)
		FROM reviews
		WHERE user_id = $1
		AND ($2::media_type_enum IS NULL OR media_type = $2::media_type_enum)
	`

	var totalItems int
	err := database.DB.QueryRow(ctx, countQuery, userID, nullableString(mediaType)).Scan(&totalItems)
	if err != nil {
		return nil, err
	}

	query := `
		SELECT 
			r.id,
			r.user_id,
			r.tmdb_id,
			r.media_type,
			r.rating,
			r.review_text,
			r.is_spoiler,
			r.created_at,
			r.updated_at,
			m.title,
			m.poster_path,
			m.backdrop_path,
			m.release_date,
			u.username AS user_name
		FROM reviews r
		JOIN media m
			ON r.tmdb_id = m.tmdb_id
			AND r.media_type = m.media_type
		JOIN users u
			ON r.user_id = u.id
		WHERE r.user_id = $1
		AND ($2::media_type_enum IS NULL OR r.media_type = $2::media_type_enum)
		ORDER BY r.updated_at DESC
		LIMIT $3 OFFSET $4
	`

	rows, err := database.DB.Query(ctx, query, userID, nullableString(mediaType), limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []models.Review
	for rows.Next() {
		var r models.Review
		var media models.MediaPreview
		err := rows.Scan(
			&r.ID,
			&r.UserID,
			&r.TMDBID,
			&r.MediaType,
			&r.Rating,
			&r.ReviewText,
			&r.IsSpoiler,
			&r.CreatedAt,
			&r.UpdatedAt,
			&media.Title,
			&media.PosterPath,
			&media.BackdropPath,
			&media.ReleaseDate,
			&r.UserName,
		)
		if err != nil {
			return nil, err
		}
		r.Media = &media
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

func GetMediaReviewsPublic(ctx context.Context, tmdbID int, mediaType string, page int, limit int) (*models.PaginatedReviews, error) {
	if page < 1 {
		page = 1
	}
	if limit <= 0 {
		limit = 12
	}
	offset := (page - 1) * limit

	countQuery := `
		SELECT COUNT(*)
		FROM reviews
		WHERE tmdb_id = $1 AND media_type = $2
	`
	var totalItems int
	err := database.DB.QueryRow(ctx, countQuery, tmdbID, mediaType).Scan(&totalItems)
	if err != nil {
		return nil, err
	}

	query := `
		SELECT 
			r.id,
			r.tmdb_id,
			r.media_type,
			r.rating,
			r.review_text,
			r.is_spoiler,
			r.created_at,
			r.updated_at,
			u.username AS user_name
		FROM reviews r
		JOIN users u ON r.user_id = u.id
		WHERE r.tmdb_id = $1 AND r.media_type = $2
		ORDER BY r.updated_at DESC
		LIMIT $3 OFFSET $4
	`

	rows, err := database.DB.Query(ctx, query, tmdbID, mediaType, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []models.Review
	for rows.Next() {
		var r models.Review
		err := rows.Scan(
			&r.ID,
			&r.TMDBID,
			&r.MediaType,
			&r.Rating,
			&r.ReviewText,
			&r.IsSpoiler,
			&r.CreatedAt,
			&r.UpdatedAt,
			&r.UserName,
		)
		if err != nil {
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
