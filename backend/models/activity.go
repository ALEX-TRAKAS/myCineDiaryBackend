package models

import "time"

type ActivityType string

const (
	ActivityAddWatchlist  ActivityType = "add_watchlist"
	ActivityRateMovie     ActivityType = "rate_movie"
	ActivityAddFavorite   ActivityType = "add_favorite"
	ActivityCompleteMovie ActivityType = "complete_movie"
	ActivityReviewMovie   ActivityType = "review_movie"
)

type Activity struct {
	ID        uint64
	UserID    uint64
	Type      ActivityType
	MovieID   *int
	SeriesID  *int
	Rating    *int
	CreatedAt time.Time
}
