package config

import (
	"fmt"
	"os"
)

// Config holds application configuration
type Config struct {
	DatabaseURL string
	GRPCPort    string
}

// Load loads configuration from environment variables
func Load() *Config {
	return &Config{
		DatabaseURL: getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/genealogy?sslmode=disable"),
		GRPCPort:    getEnv("GRPC_PORT", "50051"),
	}
}

// getEnv gets an environment variable with a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// GetDatabaseURL returns the database connection URL
func (c *Config) GetDatabaseURL() string {
	return c.DatabaseURL
}

// GetGRPCAddr returns the gRPC server address
func (c *Config) GetGRPCAddr() string {
	return fmt.Sprintf(":%s", c.GRPCPort)
}
