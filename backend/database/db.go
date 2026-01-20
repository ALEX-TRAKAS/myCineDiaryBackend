package database

import (
	"context"
	"log"
	"mycinediarybackend/config"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Connect() {
	config.Load()

	dsn := config.GetEnv("DATABASE_URL", os.Getenv("DATABASE_URL"))

	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatalf("Failed to parse DB config: %v", err)
	}

	cfg.MaxConns = 10
	cfg.MinConns = 2
	cfg.MaxConnLifetime = time.Hour

	pool, err := pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	var version string
	if err := pool.QueryRow(context.Background(), "SELECT version()").Scan(&version); err != nil {
		log.Fatalf("DB test query failed: %v", err)
	}

	log.Println("Connected to:", version)

	DB = pool
}
