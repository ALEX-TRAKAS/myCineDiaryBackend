package repositories

import (
	"context"
	"mycinediarybackend/database"
	"mycinediarybackend/models"
	"time"
)

func AddThreadPost(ctx context.Context, threadPost *models.ThreadPost) error {
	query := `
		INSERT INTO thread_posts (thread_id, user_id, Body, created_at)
		VALUES ($1, $2, $3, $4)
	`
	_, err := database.DB.Exec(
		ctx,
		query,
		threadPost.ThreadID,
		threadPost.UserID,
		threadPost.Body,
		threadPost.CreatedAt,
	)
	return err
}

func RemoveThreadPost(ctx context.Context, threadPostID uint64) error {
	query := `
		UPDATE thread_posts
		SET is_deleted = TRUE
		WHERE id = $1
	`
	_, err := database.DB.Exec(ctx, query, threadPostID)
	return err
}

func GetThreadPostsByThreadID(ctx context.Context, threadID uint64) ([]models.ThreadPost, error) {
	query := `
		SELECT id, thread_id, user_id, body, like_count, is_deleted, created_at, edited_at
		FROM thread_posts
		WHERE thread_id = $1 AND is_deleted = FALSE
		ORDER BY created_at ASC
	`
	rows, err := database.DB.Query(ctx, query, threadID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var threadPosts []models.ThreadPost
	for rows.Next() {
		var threadPost models.ThreadPost
		err := rows.Scan(
			&threadPost.ID,
			&threadPost.ThreadID,
			&threadPost.UserID,
			&threadPost.Body,
			&threadPost.LikeCount,
			&threadPost.IsDeleted,
			&threadPost.CreatedAt,
			&threadPost.EditedAt,
		)
		if err != nil {
			return nil, err
		}
		threadPosts = append(threadPosts, threadPost)
	}
	return threadPosts, nil
}

func UpdateThreadPostBody(ctx context.Context, threadPostID uint64, newBody string, editedAt time.Time) error {
	query := `
		UPDATE thread_posts
		SET body = $1, edited_at = $2
		WHERE id = $3
	`
	_, err := database.DB.Exec(ctx, query, newBody, editedAt, threadPostID)
	return err
}
