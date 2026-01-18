package repositories

import (
	"mycinediarybackend/database"
	"mycinediarybackend/models"
)

func AddUserMovie(userMovie *models.UserMovie) error {
	query := `
		INSERT INTO user_movies (user_id, tmdb_movie_id, watched_at, rating, progress)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := database.DB.Exec(
		query,
		userMovie.UserID,
		userMovie.TMDBMovieID,
		userMovie.WatchedAt,
		userMovie.Rating,
		userMovie.Progress,
	)
	return err
}

func RemoveUserMovie(userID uint64, tmdbMovieID int) error {
	query := `
		DELETE FROM user_movies
		WHERE user_id = $1 AND tmdb_movie_id = $2
	`
	_, err := database.DB.Exec(query, userID, tmdbMovieID)
	return err
}

func AddUserSeries(userSeries *models.UserSeries) error {
	query := `
		INSERT INTO user_series (user_id, tmdb_series_id, watched_at, rating, progress)
		VALUES ($1, $2, $3, $4, $5) `

	_, err := database.DB.Exec(
		query,
		userSeries.UserID,
		userSeries.TMDBSeriesID,
		userSeries.WatchedAt,
		userSeries.Rating,
		userSeries.Progress,
	)
	return err
}

func RemoveUserSeries(userID uint64, tmdbSeriesID int) error {
	query := `
		DELETE FROM user_series
		WHERE user_id = $1 AND tmdb_series_id = $2
	`
	_, err := database.DB.Exec(query, userID, tmdbSeriesID)
	return err
}

func GetUserMovies(userID uint64) ([]models.UserMovie, error) {
	query := `
		SELECT user_id, tmdb_movie_id, watched_at, rating, progress
		FROM user_movies
		WHERE user_id = $1
	`
	rows, err := database.DB.Query(query, userID)
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

	return userMovies, nil
}

func GetUserSeries(userID uint64) ([]models.UserSeries, error) {
	query := `
		SELECT user_id, tmdb_series_id, watched_at, rating, progress		
		FROM user_series
		WHERE user_id = $1
	`
	rows, err := database.DB.Query(query, userID)
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
