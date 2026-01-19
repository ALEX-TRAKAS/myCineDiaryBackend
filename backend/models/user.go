package models

import "time"

type User struct {
	ID           uint64    `db:"id" json:"id"`
	Email        string    `db:"email" json:"email"`
	Username     string    `db:"username" json:"username"`
	PasswordHash string    `db:"password_hash" json:"-"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
}
