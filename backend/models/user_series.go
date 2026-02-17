package models

import "time"

type UserSeries struct {
	UserID       uint64    `db:"user_id"`
	TMDBSeriesID int       `db:"tmdb_series_id"`
	WatchedAt    time.Time `db:"watched_at"`
	Rating       *int      `db:"rating"`
	Progress     *int      `db:"progress"`
}

type PaginatedUserSeries struct {
	Series      []UserSeries
	CurrentPage int
	TotalPages  int
	TotalItems  int
}
