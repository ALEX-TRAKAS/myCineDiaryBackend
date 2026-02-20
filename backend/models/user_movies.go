package models

import (
	"time"
)

type UserMovie struct {
	UserID      uint64    `db:"user_id"`
	TMDBMovieID int       `db:"tmdb_movie_id"`
	PosterPath  string    `db:"poster_path"`
	Title       string    `db:"title"`
	WatchedAt   time.Time `db:"watched_at"`
	Rating      *int      `db:"rating"`
	Progress    *int      `db:"progress"`
}

type PaginatedUserMovies struct {
	Movies      []UserMovie
	CurrentPage int
	TotalPages  int
	TotalItems  int
}
