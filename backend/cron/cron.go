package cron

import (
	"context"
	"time"

	"mycinediarybackend/repositories"
)

func StartTokenCleanupJob() {
	ticker := time.NewTicker(1 * time.Hour)

	go func() {
		for range ticker.C {
			repositories.CleanupExpiredTokens(context.Background())
		}
	}()
}
