package database

import (
	"context"
	"log"
	"mycinediarybackend/config"
	"os"

	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func Connect() {
	config.Load()
	conn, err := pgx.Connect(context.Background(), config.GetEnv("DATABASE_URL", os.Getenv("DATABASE_URL")))
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	var version string
	if err := conn.QueryRow(context.Background(), "SELECT version()").Scan(&version); err != nil {
		log.Fatalf("Query failed: %v", err)
	}

	log.Println("Connected to:", version)

	DB = conn
}
