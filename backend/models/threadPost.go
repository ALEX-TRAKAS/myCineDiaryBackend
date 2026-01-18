package models

import "time"

type ThreadPost struct {
	ID        uint64    `db:"id"`
	ThreadID  uint64    `db:"thread_id"`
	UserID    uint64    `db:"user_id"`
	Body      string    `db:"body"`
	LikeCount int       `db:"like_count"`
	IsDeleted bool      `db:"is_deleted"`
	CreatedAt time.Time `db:"created_at"`
	EditedAt  time.Time `db:"edited_at"`
}
