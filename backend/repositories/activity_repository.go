package repositories

import (
	"context"

	"mycinediarybackend/database"
	"mycinediarybackend/models"
)

func CreateActivity(ctx context.Context, activity *models.Activity) error {

	query := `
		INSERT INTO activities
		(id, user_id, type, movie_id, series_id, rating, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := database.DB.Exec(
		ctx,
		query,
		activity.ID,
		activity.UserID,
		activity.Type,
		activity.MovieID,
		activity.SeriesID,
		activity.Rating,
		activity.CreatedAt,
	)

	return err
}

func GetUserActivities(ctx context.Context, userID uint64, limit int) ([]models.Activity, error) {

	query := `
		SELECT
			id,
			user_id,
			type,
			movie_id,
			series_id,
			rating,
			created_at
		FROM activities
		WHERE user_id = $1
		ORDER BY created_at DESC
		LIMIT $2
	`

	rows, err := database.DB.Query(ctx, query, userID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var activities []models.Activity

	for rows.Next() {
		var a models.Activity

		err := rows.Scan(
			&a.ID,
			&a.UserID,
			&a.Type,
			&a.MovieID,
			&a.SeriesID,
			&a.Rating,
			&a.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		activities = append(activities, a)
	}

	return activities, nil
}
