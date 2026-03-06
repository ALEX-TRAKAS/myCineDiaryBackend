package repositories

import (
	"context"
	"mycinediarybackend/database"
	"mycinediarybackend/models"
)

func GetMediaByTMDBID(ctx context.Context, tmdbID int, mediaType string) (*models.Media, error) {
	query := `
		SELECT 
			m.tmdb_id,
			m.media_type,
			m.title,
			m.poster_path,
			m.backdrop_path,
			m.overview,
			m.release_date,
			m.average_rating,
			m.reviews_count
		FROM media m
		WHERE m.tmdb_id = $1 AND m.media_type = $2s
	`

	var m models.Media

	err := database.DB.QueryRow(ctx, query, tmdbID, mediaType).Scan(
		&m.TMDBID,
		&m.MediaType,
		&m.Title,
		&m.PosterPath,
		&m.BackdropPath,
		&m.Overview,
		&m.ReleaseDate,
		&m.AverageRating,
		&m.ReviewsCount,
	)

	if err != nil {
		return nil, err
	}

	return &m, nil
}
