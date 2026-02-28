package models

import "time"

type Review struct {
	ID         uint64        `json:"id"`
	UserID     uint64        `json:"user_id"`
	UserName   string        `json:"user_name,omitempty"`
	TMDBID     int           `json:"tmdb_id"`
	MediaType  MediaType     `json:"media_type"`
	Rating     int           `json:"rating"`
	ReviewText string        `json:"review_text,omitempty"`
	IsSpoiler  bool          `json:"is_spoiler"`
	CreatedAt  time.Time     `json:"created_at"`
	UpdatedAt  time.Time     `json:"updated_at"`
	Media      *MediaPreview `json:"media,omitempty"`
}

type PaginatedReviews struct {
	Reviews     []Review `json:"reviews"`
	CurrentPage int      `json:"current_page"`
	TotalPages  int      `json:"total_pages"`
	TotalItems  int      `json:"total_items"`
}

type MediaPreview struct {
	Title        string     `json:"title"`
	PosterPath   string     `json:"poster_path"`
	BackdropPath string     `json:"backdrop_path"`
	ReleaseDate  *time.Time `json:"release_date,omitempty"`
}
