package models

import "time"

type RefreshToken struct {
	ID        uint64    `db:"id"`
	UserID    uint64    `db:"user_id"`
	TokenSHA  string    `db:"token_sha"`
	TokenHash string    `db:"token_hash"`
	FamilyID  string    `db:"family_id"`
	Revoked   bool      `db:"revoked"`
	ExpiresAt time.Time `db:"expires_at"`
}
