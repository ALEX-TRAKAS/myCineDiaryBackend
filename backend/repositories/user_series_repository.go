package repositories

import (
	"context"
	"mycinediarybackend/database"
	"mycinediarybackend/models"
)

func AddUserSeries(ctx context.Context, userSeries *models.UserSeries) error {
	query := `
		INSERT INTO user_series (user_id, tmdb_series_id, watched_at, rating, progress)
		VALUES ($1, $2, $3, $4, $5) `

	_, err := database.DB.Exec(
		ctx,
		query,
		userSeries.UserID,
		userSeries.TMDBSeriesID,
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

func GetUserSeries(ctx context.Context, userID uint64) ([]models.UserSeries, error) {
	query := `
		SELECT user_id, tmdb_series_id, watched_at, rating, progress		
		FROM user_series
		WHERE user_id = $1
	`
	rows, err := database.DB.Query(ctx, query, userID)
	if err != nil {
		return nil, err

	}
	defer rows.Close()
	var userSeriesList []models.UserSeries
	for rows.Next() {
		var userSeries models.UserSeries
		err := rows.Scan(
			&userSeries.UserID,
			&userSeries.TMDBSeriesID,
			&userSeries.WatchedAt,
			&userSeries.Rating,
			&userSeries.Progress,
		)
		if err != nil {
			return nil, err
		}
		userSeriesList = append(userSeriesList, userSeries)
	}

	return userSeriesList, nil
}
