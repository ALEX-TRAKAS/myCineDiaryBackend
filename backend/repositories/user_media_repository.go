package repositories

import (
	"context"
	"mycinediarybackend/database"
	"mycinediarybackend/models"
)

func AddUserMedia(ctx context.Context, userMedia *models.UserMedia) error {
	tx, err := database.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	mediaQuery := `
		INSERT INTO media (
			tmdb_id, media_type, title, poster_path,
			backdrop_path, overview, release_date, created_at, updated_at
		)
		VALUES ($1,$2,$3,$4,$5,$6,$7,NOW(),NOW())
		ON CONFLICT (tmdb_id, media_type) DO UPDATE
		SET
			title = EXCLUDED.title,
			poster_path = EXCLUDED.poster_path,
			backdrop_path = EXCLUDED.backdrop_path,
			overview = EXCLUDED.overview,
			release_date = EXCLUDED.release_date,
			updated_at = NOW()
	`

	_, err = tx.Exec(
		ctx,
		mediaQuery,
		userMedia.TMDBID,
		userMedia.MediaType,
		userMedia.Title,
		userMedia.PosterPath,
		userMedia.BackdropPath,
		userMedia.Overview,
		userMedia.ReleaseDate,
	)
	if err != nil {
		return err
	}

	userMediaQuery := `
		INSERT INTO user_media (
			user_id, tmdb_id, media_type,
			rating, progress, notes,
			status, is_favorite, created_at, updated_at
		)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,NOW(),NOW())
		ON CONFLICT (user_id, tmdb_id, media_type) DO UPDATE
		SET
			rating = EXCLUDED.rating,
			progress = EXCLUDED.progress,
			notes = EXCLUDED.notes,
			status = EXCLUDED.status,
			is_favorite = EXCLUDED.is_favorite,
			updated_at = NOW()
	`

	_, err = tx.Exec(
		ctx,
		userMediaQuery,
		userMedia.UserID,
		userMedia.TMDBID,
		userMedia.MediaType,
		userMedia.Rating,
		userMedia.Progress,
		userMedia.Notes,
		userMedia.Status,
		userMedia.IsFavorite,
	)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
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
		FROM user_media um
		WHERE um.user_id = $1 
		AND ($2::media_type_enum IS NULL OR um.media_type = $2::media_type_enum)
	`
	err := database.DB.QueryRow(ctx, countQuery, userID, nullableString(mediaType)).Scan(&totalItems)
	if err != nil {
		return nil, err
	}

	query := `
		SELECT 
			um.user_id,
			um.tmdb_id,
			um.media_type,
			um.rating,
			um.progress,
			um.notes,
			um.status,
			um.is_favorite,
			um.created_at,
			um.updated_at,
			m.title,
			m.poster_path,
			m.backdrop_path,
			m.overview,
			m.release_date

		FROM user_media um
		JOIN media m
			ON m.tmdb_id = um.tmdb_id
			AND m.media_type = um.media_type

		WHERE um.user_id = $1
		AND ($2::media_type_enum IS NULL OR um.media_type = $2::media_type_enum)

		ORDER BY um.updated_at DESC
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

			&m.Title,
			&m.PosterPath,
			&m.BackdropPath,
			&m.Overview,
			&m.ReleaseDate,
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

func GetUserMediaByTMDBID(ctx context.Context, userID uint64, tmdbID int, mediaType string) (*models.UserMedia, error) {
	query := `
		SELECT 
			um.user_id,
			um.tmdb_id,
			um.media_type,
			um.rating,
			um.progress,
			um.notes,
			um.status,
			um.is_favorite,
			um.created_at,
			um.updated_at,
			m.title,
			m.poster_path,
			m.backdrop_path,
			m.overview,
			m.release_date

		FROM user_media um
		JOIN media m
			ON m.tmdb_id = um.tmdb_id
			AND m.media_type = um.media_type

		WHERE um.user_id = $1 AND um.tmdb_id = $2 AND um.media_type = $3
	`

	var m models.UserMedia

	err := database.DB.QueryRow(ctx, query, userID, tmdbID, mediaType).Scan(
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

		&m.Title,
		&m.PosterPath,
		&m.BackdropPath,
		&m.Overview,
		&m.ReleaseDate)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

func nullableString(s string) interface{} {
	if s == "" {
		return nil
	}
	return s
}
