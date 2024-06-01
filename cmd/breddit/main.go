package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/Potagashev/breddit/internal/config"
	"github.com/jackc/pgx/v5"
)


const (
	envLocal = "local"
	envDev = "dev"
	envProd = "prod"
)


func main() {
	cfg := config.MustLoad()
	
	logger := setupLogger(cfg.Env)

	initTables(cfg.DbUrl)
	logger.Info("db tables created")

	// TODO init router: chi, "chi render"

	// TODO run server
}

func setupLogger(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case envLocal:
		logger = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		logger = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return logger
}

func initTables(dbUrl string) error {
	log.Printf("Connecting to database: %s", dbUrl)
	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
		
	_, err = conn.Exec(
		context.Background(),
		`
		CREATE TABLE IF NOT EXISTS threads (
			id UUID PRIMARY KEY,
			title TEXT NOT NULL,
			text TEXT NOT NULL,
			created_at TIMESTAMPTZ NOT NULL,
			updated_at TIMESTAMPTZ NOT NULL
		)
		`,
	)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}