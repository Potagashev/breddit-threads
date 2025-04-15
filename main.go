package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Potagashev/breddit_threads/internal/auth"
	"github.com/Potagashev/breddit_threads/internal/config"
	"github.com/Potagashev/breddit_threads/internal/router"
	"github.com/Potagashev/breddit_threads/internal/threads"
	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server.
// @host localhost:8080
// @BasePath /api/v1
func main() {
	cfg, _ := config.LoadConfig()

	dbConn, err := pgx.Connect(context.Background(), cfg.DbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbConn.Close(context.Background())

	conn, err := grpc.NewClient("app-auth:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Fatalf("Failed to connect to gRPC server: %v", err)
    }
    defer conn.Close()

	authService := auth.NewAuthService(conn)
	threadsRepository := threads.NewThreadRepository(dbConn)
	threadsService := threads.NewThreadService(threadsRepository)
	threadsHandler := threads.NewThreadHandler(threadsService, authService, cfg)

	r := router.NewRouter(threadsHandler)

	r.Run(fmt.Sprintf(":%s", cfg.AppPort))
}
