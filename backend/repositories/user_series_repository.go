package repositories

import (
	"context"
	"mycinediarybackend/database"
	"mycinediarybackend/models"
)

func AddUserSeries(ctx context.Context, userSeries *models.UserSeries) error {
	query := `
		INSERT INTO user_series (user_id, tmdb_series_id, poster_path, title, watched_at, rating, progress)
		VALUES ($1, $2, $3, $4, $5, $6, $7) `

	_, err := database.DB.Exec(
		ctx,
		query,
		userSeries.UserID,
		userSeries.TMDBSeriesID,
		userSeries.PosterPath,
		userSeries.Title,
		userSeries.WatchedAt,
		userSeries.Rating,
		userSeries.Progress,
	)
	return err
}

func RemoveUserSeries(ctx context.Context, userID uint64, tmdbSeriesID int) error {
	query := `
		DELETE FROM user_series
		WHERE user_id = $1 AND tmdb_series_id = $2
	`
	_, err := database.DB.Exec(ctx, query, userID, tmdbSeriesID)
	return err
}

func GetUserSeries(ctx context.Context, userID uint64, page int, limit int) (*models.PaginatedUserSeries, error) {

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
		FROM user_series
		WHERE user_id = $1
	`
	err := database.DB.QueryRow(ctx, countQuery, userID).Scan(&totalItems)
	if err != nil {
		return nil, err
	}

	query := `
		SELECT user_id, tmdb_series_id, poster_path, title, watched_at, rating, progress
		FROM user_series
		WHERE user_id = $1
		ORDER BY watched_at DESC
		LIMIT $2 OFFSET $3
	`

	rows, err := database.DB.Query(ctx, query, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userSeries []models.UserSeries

	for rows.Next() {
		var userSeriesItem models.UserSeries
		err := rows.Scan(
			&userSeriesItem.UserID,
			&userSeriesItem.TMDBSeriesID,
			&userSeriesItem.PosterPath,
			&userSeriesItem.Title,
			&userSeriesItem.WatchedAt,
			&userSeriesItem.Rating,
			&userSeriesItem.Progress,
		)
		if err != nil {
			return nil, err
		}
		userSeries = append(userSeries, userSeriesItem)
	}

	totalPages := (totalItems + limit - 1) / limit

	return &models.PaginatedUserSeries{
		Series:      userSeries,
		CurrentPage: page,
		TotalPages:  totalPages,
		TotalItems:  totalItems,
	}, nil
}
