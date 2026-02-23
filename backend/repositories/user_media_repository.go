package repositories

import (
	"context"
	"mycinediarybackend/database"
	"mycinediarybackend/models"
)

func AddUserMedia(ctx context.Context, userMedia *models.UserMedia) error {
	query := `
		INSERT INTO user_media (
			user_id, tmdb_id, media_type, rating, progress, notes,
			status, is_favorite, created_at, updated_at
		)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,NOW(),NOW())
		ON CONFLICT (user_id, tmdb_id, media_type) DO UPDATE
		SET updated_at = NOW(),
		    rating = EXCLUDED.rating,
		    progress = EXCLUDED.progress,
		    notes = EXCLUDED.notes,
		    status = EXCLUDED.status,
		    is_favorite = EXCLUDED.is_favorite
	`
	_, err := database.DB.Exec(
		ctx,
		query,
		userMedia.UserID,
		userMedia.TMDBID,
		userMedia.MediaType,
		userMedia.Rating,
		userMedia.Progress,
		userMedia.Notes,
		userMedia.Status,
		userMedia.IsFavorite,
	)
	return err
}

func RemoveUserMedia(ctx context.Context, userID uint64, tmdbID int, mediaType string) error {
	query := `
		DELETE FROM user_media
		WHERE user_id = $1 AND tmdb_id = $2 AND media_type = $3
	`
	_, err := database.DB.Exec(ctx, query, userID, tmdbID, mediaType)
	return err
}

func GetUserMedia(ctx context.Context, userID uint64, mediaType string, page int, limit int) (*models.PaginatedUserMedia, error) {
	if page < 1 {
		page = 1
	}
	if limit <= 0 {
		limit = 12
	}

	offset := (page - 1) * limit

	var totalItems int
	countQuery := `
		SELECT COUNT(*)
		FROM user_media
		WHERE user_id = $1 AND ($2 IS NULL OR media_type = $2)
	`
	err := database.DB.QueryRow(ctx, countQuery, userID, nullableString(mediaType)).Scan(&totalItems)
	if err != nil {
		return nil, err
	}

	query := `
		SELECT 
			user_id, tmdb_id, media_type, rating, progress, notes, 
			status, is_favorite, created_at, updated_at
		FROM user_media
		WHERE user_id = $1 AND ($2 IS NULL OR media_type = $2)
		ORDER BY updated_at DESC
		LIMIT $3 OFFSET $4
	`
	rows, err := database.DB.Query(ctx, query, userID, nullableString(mediaType), limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userMediaList []models.UserMedia
	for rows.Next() {
		var m models.UserMedia
		err := rows.Scan(
			&m.UserID,
			&m.TMDBID,
			&m.MediaType,
			&m.Rating,
			&m.Progress,
			&m.Notes,
			&m.Status,
			&m.IsFavorite,
			&m.CreatedAt,
			&m.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		userMediaList = append(userMediaList, m)
	}

	totalPages := (totalItems + limit - 1) / limit

	return &models.PaginatedUserMedia{
		Items:       userMediaList,
		CurrentPage: page,
		TotalPages:  totalPages,
		TotalItems:  totalItems,
	}, nil
}

func nullableString(s string) interface{} {
	if s == "" {
		return nil
	}
	return s
}
