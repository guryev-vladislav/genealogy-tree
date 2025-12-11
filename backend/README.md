# Backend

This is the backend service for the genealogy-tree project, built with Go, gRPC, and PostgreSQL.

## Architecture

The backend follows a clean architecture pattern with the following structure:

- `cmd/` - Application entry point (main.go)
- `config/` - Configuration management
- `internal/model/` - Domain models
- `internal/repository/` - Data access layer
- `sql/` - Database initialization scripts

## Prerequisites

- Go 1.21 or higher
- PostgreSQL 15 or higher
- Docker and Docker Compose (for local development)

## Getting Started

### 1. Start the PostgreSQL database

```bash
# From the root of the repository
docker-compose up -d
```

This will start a PostgreSQL instance with the database initialized using `sql/init.sql`.

### 2. Build the application

```bash
cd backend
go mod tidy
go build -o genealogy-server ./cmd/main.go
```

### 3. Run the application

```bash
# Set environment variables (optional, defaults are provided)
export DATABASE_URL="postgres://postgres:postgres@localhost:5432/genealogy?sslmode=disable"
export GRPC_PORT="50051"

# Run the server
./genealogy-server
```

## Environment Variables

- `DATABASE_URL` - PostgreSQL connection string (default: `postgres://postgres:postgres@localhost:5432/genealogy?sslmode=disable`)
- `GRPC_PORT` - gRPC server port (default: `50051`)

## Database Schema

### persons table

- `id` (UUID) - Primary key
- `name` (VARCHAR) - Person's name
- `dates` (VARCHAR) - Birth/death dates or date range
- `created_at` (TIMESTAMP) - Record creation timestamp
- `updated_at` (TIMESTAMP) - Record update timestamp

## Development

The application uses:

- [pgx](https://github.com/jackc/pgx) - PostgreSQL driver and toolkit
- [gRPC](https://grpc.io/) - RPC framework
- [uuid](https://github.com/google/uuid) - UUID generation

## Next Steps

- Define Protocol Buffer schemas for gRPC services
- Implement gRPC service handlers
- Add business logic layer
- Implement authentication and authorization
