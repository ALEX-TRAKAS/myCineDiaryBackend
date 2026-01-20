package repositories

import (
	"context"
	"mycinediarybackend/database"
	"mycinediarybackend/models"
	"time"
)

func SaveRefreshToken(ctx context.Context, userID uint64, sha string, hash string, familyID string, ip string, deviceID string, expiresAt time.Time) error {
	query := `
        INSERT INTO refresh_tokens (user_id, token_sha, token_hash, family_id, created_ip, device_id, expires_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
    `
	_, err := database.DB.Exec(ctx, query, userID, sha, hash, familyID, ip, deviceID, expiresAt)
	return err
}

func GetRefreshTokenBySHA(ctx context.Context, sha string) (*models.RefreshToken, error) {
	query := `
        SELECT id, user_id, token_sha, token_hash, family_id, revoked, expires_at
        FROM refresh_tokens
        WHERE token_sha = $1
          AND revoked = FALSE
          AND expires_at > NOW()
    `
	var rt models.RefreshToken
	err := database.DB.QueryRow(ctx, query, sha).Scan(
		&rt.ID, &rt.UserID, &rt.TokenSHA, &rt.TokenHash,
		&rt.FamilyID, &rt.Revoked, &rt.ExpiresAt,
	)
	if err != nil {
		return nil, err
	}
	return &rt, nil
}

func RevokeRefreshTokenByID(ctx context.Context, id uint64) error {
	query := `UPDATE refresh_tokens SET revoked = TRUE WHERE id = $1`
	_, err := database.DB.Exec(ctx, query, id)
	return err
}

func RevokeFamilyTokensByFamilyID(ctx context.Context, familyID string) error {
	query := `UPDATE refresh_tokens SET revoked = TRUE WHERE family_id = $1`
	_, err := database.DB.Exec(ctx, query, familyID)
	return err
}

func RevokeAllRefreshTokensForUser(ctx context.Context, userID uint64) error {
	query := `UPDATE refresh_tokens SET revoked = true WHERE user_id = $1`
	_, err := database.DB.Exec(ctx, query, userID)
	return err
}

func CleanupExpiredTokens(ctx context.Context) error {
	query := `DELETE FROM refresh_tokens WHERE expires_at < NOW()`
	_, err := database.DB.Exec(ctx, query)
	return err
}

func UpdateRefreshTokenUsage(ctx context.Context, id uint64, ip string) error {
	query := `
        UPDATE refresh_tokens
        SET last_used_ip = $1,
            last_used_at = NOW()
        WHERE id = $2
    `
	_, err := database.DB.Exec(ctx, query, ip, id)
	return err
}
