package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/ScMofeoluwa/ingestor/internal/config"
)

func SetupDB(cfg config.Config) *Queries {
	conn, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("error: %s\n", err)
	}
	return New(conn)
}
