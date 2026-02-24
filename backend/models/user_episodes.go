package models

import "time"

type UserEpisode struct {
	UserID        uint64     `json:"user_id"`
	TMDBID        int        `json:"tmdb_id"`
	MediaType     MediaType  `json:"media_type"`
	SeasonNumber  int        `json:"season_number"`
	EpisodeNumber int        `json:"episode_number"`
	Watched       bool       `json:"watched"`
	WatchedAt     *time.Time `json:"watched_at,omitempty"`
	CreatedAt     time.Time  `json:"created_at"`
}

type PaginatedUserEpisodes struct {
	Items       []UserEpisode `json:"items"`
	CurrentPage int           `json:"current_page"`
	TotalPages  int           `json:"total_pages"`
	TotalItems  int           `json:"total_items"`
}
