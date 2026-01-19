package repositories

import (
	"mycinediarybackend/database"
	"mycinediarybackend/models"
)

func CreateThread(thread *models.Thread) error {
	query := `
		INSERT INTO threads (title, created_at)
		VALUES ($1, $2)
	`
	_, err := database.DB.Exec(
		query,
		thread.Title,
		thread.CreatedAt,
	)
	return err
}

func GetAllThreads() ([]models.Thread, error) {
	query := `
		SELECT id, title, created_at
		FROM threads
		ORDER BY created_at DESC
	`
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var threads []models.Thread
	for rows.Next() {
		var thread models.Thread
		err := rows.Scan(
			&thread.ID,
			&thread.Title,
			&thread.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		threads = append(threads, thread)
	}
	return threads, nil
}

func GetThreadByID(threadID string) (*models.Thread, error) {
	query := `
		SELECT id, title, created_at
		FROM threads
		WHERE id = $1
	`
	var thread models.Thread
	err := database.DB.QueryRow(query, threadID).Scan(
		&thread.ID,
		&thread.Title,
		&thread.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &thread, nil
}

func DeleteThread(threadID string) error {
	query := `
		DELETE FROM threads
		WHERE id = $1
	`
	_, err := database.DB.Exec(query, threadID)
	return err
}
