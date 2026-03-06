package models

import "time"

type UserMedia struct {
	UserID       uint64      `json:"user_id"`
	TMDBID       int         `json:"tmdb_id"`
	MediaType    MediaType   `json:"media_type"`
	Rating       *int        `json:"rating,omitempty"`
	Progress     *int        `json:"progress,omitempty"`
	Notes        string      `json:"notes,omitempty"`
	Status       MediaStatus `json:"status"`
	IsFavorite   bool        `json:"is_favorite"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
	Title        string      `json:"title"`
	PosterPath   string      `json:"poster_path"`
	BackdropPath string      `json:"backdrop_path"`
	Overview     string      `json:"overview"`
	ReleaseDate  *time.Time  `json:"release_date,omitempty"`
	Genres       []Genre     `json:"genres" gorm:"type:jsonb"`
}

type PaginatedUserMedia struct {
	Items       []UserMedia `json:"items"`
	CurrentPage int         `json:"current_page"`
	TotalPages  int         `json:"total_pages"`
	TotalItems  int         `json:"total_items"`
}
