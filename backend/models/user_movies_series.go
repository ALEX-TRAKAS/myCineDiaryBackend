package models

import "time"

type UserMovie struct {
	UserID      uint64    `db:"user_id"`
	TMDBMovieID int       `db:"tmdb_movie_id"`
	WatchedAt   time.Time `db:"watched_at"`
	Rating      *int      `db:"rating"`
	Progress    *int      `db:"progress"`
}

type UserSeries struct {
	UserID       uint64    `db:"user_id"`
	TMDBSeriesID int       `db:"tmdb_series_id"`
	WatchedAt    time.Time `db:"watched_at"`
	Rating       *int      `db:"rating"`
	Progress     *int      `db:"progress"`
}
