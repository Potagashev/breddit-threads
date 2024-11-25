package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"github.com/Potagashev/breddit/internal/config"
	"github.com/jackc/pgx/v5"
	"github.com/Potagashev/breddit/internal/threads"
	"github.com/Potagashev/breddit/internal/router"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server.
// @host localhost:8080
// @BasePath /api/v1
func main() {
	cfg, _ := config.LoadConfig()
	
	err := initTables(cfg.DbUrl)
	if err != nil {
		log.Println("db was NOT created")
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
	}
	log.Println("db tables created")

	conn, err := pgx.Connect(context.Background(), cfg.DbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	threads_repository := threads.NewThreadRepository(conn)
	threads_service := threads.NewThreadService(threads_repository)
	
	r := router.NewRouter(threads_service)
	
	r.Run(fmt.Sprintf(":%s", cfg.AppPort))
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
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			title TEXT NOT NULL,
			text TEXT NOT NULL,
			created_at TIMESTAMPTZ NOT NULL default now(),
			updated_at TIMESTAMPTZ NOT NULL default now()
		)
		`,
	)
	if err != nil {
		return err
	}
	log.Println("Tables has been created")

	return nil
}