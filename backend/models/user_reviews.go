package models

import "time"

type Review struct {
	ID         uint64    `json:"id"`
	UserID     uint64    `json:"user_id"`
	TMDBID     int       `json:"tmdb_id"`
	MediaType  string    `json:"media_type"` // "movie" or "tv"
	Rating     int       `json:"rating"`
	ReviewText string    `json:"review_text,omitempty"`
	IsSpoiler  bool      `json:"is_spoiler"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type PaginatedReviews struct {
	Reviews     []Review `json:"reviews"`
	CurrentPage int      `json:"current_page"`
	TotalPages  int      `json:"total_pages"`
	TotalItems  int      `json:"total_items"`
}
