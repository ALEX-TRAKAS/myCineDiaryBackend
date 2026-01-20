package services

import (
	"context"
	"errors"
	"mycinediarybackend/models"
	"mycinediarybackend/repositories"
	"mycinediarybackend/utils"
	"time"
)

func CreateRefreshToken(ctx context.Context, userID uint64, familyID string, ip string, deviceID string) (string, error) {
	token, _ := utils.GenerateRefreshToken()
	hash, _ := utils.HashToken(token)
	sha := utils.SHA256(token)

	err := repositories.SaveRefreshToken(ctx, userID, sha, hash, familyID, ip, deviceID, time.Now().Add(7*24*time.Hour))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ValidateRefreshToken(ctx context.Context, token string) (*models.RefreshToken, error) {
	sha := utils.SHA256(token)
	stored, err := repositories.GetRefreshTokenBySHA(ctx, sha)
	if err != nil {
		return nil, err
	}

	if !utils.CompareHashAndToken(stored.TokenHash, token) {
		return nil, errors.New("invalid refresh token")
	}
	return stored, nil
}

func LogoutAll(ctx context.Context, userID uint64) error {
	return repositories.RevokeAllRefreshTokensForUser(ctx, userID)
}
