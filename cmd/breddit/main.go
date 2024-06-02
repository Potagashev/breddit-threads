package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"github.com/Potagashev/breddit/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/Potagashev/breddit/internal/threads"
	"github.com/Potagashev/breddit/internal/router"
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

	conn, err := pgx.Connect(context.Background(), cfg.DbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	threads_repository := threads.NewThreadRepository(conn)
	threads_service := threads.NewThreadService(threads_repository)
	
	r := router.NewRouter(threads_service)
	
	r.Run("localhost:8080")
}

func getHelloWorld(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World!",
	})
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
			created_at TIMESTAMPTZ NOT NULL default now(),
			updated_at TIMESTAMPTZ NOT NULL default now()
		)
		`,
	)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}