package models

import "time"

type User struct {
	ID        uint64    `db:"id"`
	Email     string    `db:"email"`
	Username  string    `db:"username"`
	CreatedAt time.Time `db:"created_at"`
}
