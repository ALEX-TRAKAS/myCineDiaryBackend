package models

import "time"

type RefreshToken struct {
	ID         uint64    `db:"id"`
	UserID     uint64    `db:"user_id"`
	TokenSHA   string    `db:"token_sha"`
	TokenHash  string    `db:"token_hash"`
	FamilyID   string    `db:"family_id"`
	DeviceID   string    `db:"device_id"`
	CreatedIP  string    `db:"created_ip"`
	LastUsedIP string    `db:"last_used_ip"`
	LastUsedAt time.Time `db:"last_used_at"`
	Revoked    bool      `db:"revoked"`
	ExpiresAt  time.Time `db:"expires_at"`
	CreatedAt  time.Time `db:"created_at"`
}
