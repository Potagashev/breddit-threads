package main

import (
	"context"
	"fmt"
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
	
	conn, err := pgx.Connect(context.Background(), cfg.DbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	threads_repository := threads.NewThreadRepository(conn)
	threads_service := threads.NewThreadService(threads_repository)
	threads_handler := threads.NewThreadHandler(threads_service, cfg)

	r := router.NewRouter(threads_handler)
	
	r.Run(fmt.Sprintf(":%s", cfg.AppPort))
}
