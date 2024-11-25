package config

import (
	"fmt"
	"os"
)

type Config struct {
	AppPort string
	
	DbName string
	DbUser string
	DbPassword string
	DbHost string
	DbPort string
	DbUrl string
}

func LoadConfig() (*Config, error) {
	dbName := mustGetEnv("POSTGRES_DB")
	dbUser := mustGetEnv("POSTGRES_USER")
	dbPassword := mustGetEnv("POSTGRES_PASSWORD")
	dbHost := mustGetEnv("DATABASE_HOST")
	dbPort := mustGetEnv("DATABASE_PORT")
	dbUrl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	appPort := getEnv("APP_PORT", "8080")

	return &Config{
		DbName: dbName,
		DbUser: dbUser,
		DbPassword: dbPassword,
		DbHost: dbHost,
		DbPort: dbPort,
		DbUrl: dbUrl,
		
		AppPort: appPort,
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func mustGetEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	fmt.Fprintf(os.Stderr, "Fatal error: environment variable %s is not set\n", key)
	os.Exit(1)
	return ""
}