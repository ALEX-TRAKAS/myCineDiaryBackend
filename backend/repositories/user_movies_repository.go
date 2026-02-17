package repositories

import (
	"context"
	"mycinediarybackend/database"
	"mycinediarybackend/models"
)

func AddUserMovie(ctx context.Context, userMovie *models.UserMovie) error {
	query := `
		INSERT INTO user_movies (user_id, tmdb_movie_id, watched_at, rating, progress)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := database.DB.Exec(
		ctx,
		query,
		userMovie.UserID,
		userMovie.TMDBMovieID,
		userMovie.WatchedAt,
		userMovie.Rating,
		userMovie.Progress,
	)
	return err
}

func RemoveUserMovie(ctx context.Context, userID uint64, tmdbMovieID int) error {
	query := `
		DELETE FROM user_movies
		WHERE user_id = $1 AND tmdb_movie_id = $2
	`
	_, err := database.DB.Exec(ctx, query, userID, tmdbMovieID)
	return err
}

func GetUserMovies(ctx context.Context, userID uint64, page int, limit int) (*models.PaginatedUserMovies, error) {

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
		FROM user_movies
		WHERE user_id = $1
	`
	err := database.DB.QueryRow(ctx, countQuery, userID).Scan(&totalItems)
	if err != nil {
		return nil, err
	}

	query := `
		SELECT user_id, tmdb_movie_id, watched_at, rating, progress
		FROM user_movies
		WHERE user_id = $1
		ORDER BY watched_at DESC
		LIMIT $2 OFFSET $3
	`

	rows, err := database.DB.Query(ctx, query, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userMovies []models.UserMovie

	for rows.Next() {
		var userMovie models.UserMovie
		err := rows.Scan(
			&userMovie.UserID,
			&userMovie.TMDBMovieID,
			&userMovie.WatchedAt,
			&userMovie.Rating,
			&userMovie.Progress,
		)
		if err != nil {
			return nil, err
		}
		userMovies = append(userMovies, userMovie)
	}

	totalPages := (totalItems + limit - 1) / limit

	return &models.PaginatedUserMovies{
		Movies:      userMovies,
		CurrentPage: page,
		TotalPages:  totalPages,
		TotalItems:  totalItems,
	}, nil
}
