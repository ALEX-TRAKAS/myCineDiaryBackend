package models

import "time"

type MediaType string

const (
	MediaTypeMovie MediaType = "movie"
	MediaTypeTV    MediaType = "tv"
)

type Media struct {
	ID            uint64     `json:"id"`
	TMDBID        int        `json:"tmdb_id"`
	MediaType     MediaType  `json:"media_type"`
	Title         string     `json:"title"`
	Overview      string     `json:"overview"`
	PosterPath    string     `json:"poster_path"`
	BackdropPath  string     `json:"backdrop_path"`
	ReleaseDate   *time.Time `json:"release_date,omitempty"`
	Popularity    float64    `json:"popularity"`
	VoteAverage   float64    `json:"vote_average"`
	AverageRating float64    `json:"average_rating" gorm:"not null;default:0"`
	ReviewsCount  int        `json:"reviews_count" gorm:"not null;default:0"`

	CachedAt time.Time `json:"cached_at"`
}

type MediaStatus string

const (
	StatusWatchlist  MediaStatus = "watchlist"
	StatusInProgress MediaStatus = "in_progress"
	StatusCompleted  MediaStatus = "completed"
)
