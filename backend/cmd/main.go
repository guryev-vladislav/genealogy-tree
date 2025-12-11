package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/guryev-vladislav/genealogy-tree/backend/config"
	"github.com/guryev-vladislav/genealogy-tree/backend/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Create context for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Initialize database connection pool
	dbPool, err := initializeDBPool(ctx, cfg.GetDatabaseURL())
	if err != nil {
		log.Fatalf("Failed to initialize database pool: %v", err)
	}
	defer dbPool.Close()

	log.Println("Database connection pool initialized successfully")

	// Initialize repository
	_ = repository.NewPersonRepositoryPGX(dbPool)

	// Initialize gRPC server
	grpcServer := grpc.NewServer()

	// TODO: Register gRPC services here when proto files are defined
	// Example: pb.RegisterPersonServiceServer(grpcServer, personService)

	// Start gRPC server
	listener, err := net.Listen("tcp", cfg.GetGRPCAddr())
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", cfg.GetGRPCAddr(), err)
	}

	log.Printf("gRPC server starting on %s", cfg.GetGRPCAddr())

	// Run gRPC server in a goroutine
	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Graceful shutdown
	grpcServer.GracefulStop()
	log.Println("Server stopped")
}

// initializeDBPool initializes and returns a pgxpool connection pool
func initializeDBPool(ctx context.Context, databaseURL string) (*pgxpool.Pool, error) {
	// Configure pool
	poolConfig, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		return nil, fmt.Errorf("unable to parse database URL: %w", err)
	}

	// Set pool configuration
	poolConfig.MaxConns = 25
	poolConfig.MinConns = 5
	poolConfig.MaxConnLifetime = time.Hour
	poolConfig.MaxConnIdleTime = 30 * time.Minute

	// Create pool
	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}

	// Test connection
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("unable to ping database: %w", err)
	}

	return pool, nil
}
