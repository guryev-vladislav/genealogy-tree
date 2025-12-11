package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/guryev-vladislav/genealogy-tree/backend/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

// PersonRepository defines the interface for person data access operations
type PersonRepository interface {
	Create(ctx context.Context, person *model.Person) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.Person, error)
	List(ctx context.Context) ([]*model.Person, error)
	Update(ctx context.Context, person *model.Person) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// PersonRepositoryPGX implements PersonRepository using pgx
type PersonRepositoryPGX struct {
	pool *pgxpool.Pool
}

// NewPersonRepositoryPGX creates a new PersonRepositoryPGX instance
func NewPersonRepositoryPGX(pool *pgxpool.Pool) PersonRepository {
	return &PersonRepositoryPGX{pool: pool}
}

// Create inserts a new person into the database
func (r *PersonRepositoryPGX) Create(ctx context.Context, person *model.Person) error {
	query := `INSERT INTO persons (id, name, dates) VALUES ($1, $2, $3)`
	_, err := r.pool.Exec(ctx, query, person.ID, person.Name, person.Dates)
	return err
}

// GetByID retrieves a person by ID from the database
func (r *PersonRepositoryPGX) GetByID(ctx context.Context, id uuid.UUID) (*model.Person, error) {
	query := `SELECT id, name, dates FROM persons WHERE id = $1`
	var person model.Person
	err := r.pool.QueryRow(ctx, query, id).Scan(&person.ID, &person.Name, &person.Dates)
	if err != nil {
		return nil, err
	}
	return &person, nil
}

// List retrieves all persons from the database
func (r *PersonRepositoryPGX) List(ctx context.Context) ([]*model.Person, error) {
	query := `SELECT id, name, dates FROM persons ORDER BY name`
	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var persons []*model.Person
	for rows.Next() {
		var person model.Person
		if err := rows.Scan(&person.ID, &person.Name, &person.Dates); err != nil {
			return nil, err
		}
		persons = append(persons, &person)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return persons, nil
}

// Update updates an existing person in the database
func (r *PersonRepositoryPGX) Update(ctx context.Context, person *model.Person) error {
	query := `UPDATE persons SET name = $2, dates = $3 WHERE id = $1`
	_, err := r.pool.Exec(ctx, query, person.ID, person.Name, person.Dates)
	return err
}

// Delete removes a person from the database
func (r *PersonRepositoryPGX) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM persons WHERE id = $1`
	_, err := r.pool.Exec(ctx, query, id)
	return err
}
