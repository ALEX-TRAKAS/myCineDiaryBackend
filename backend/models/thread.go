package models

import "time"

type Thread struct {
	ID         uint64    `db:"id"`
	UserID     uint64    `db:"user_id"`
	Title      string    `db:"title"`
	Body       string    `db:"body"`
	ReplyCount int       `db:"reply_count"`
	LikeCount  int       `db:"like_count"`
	IsLocked   bool      `db:"is_locked"`
	IsDeleted  bool      `db:"is_deleted"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
