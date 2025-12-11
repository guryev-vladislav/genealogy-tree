package model

import "github.com/google/uuid"

// Person represents a person in the genealogy tree
type Person struct {
	ID    uuid.UUID `json:"id" db:"id"`
	Name  string    `json:"name" db:"name"`
	Dates string    `json:"dates" db:"dates"`
}
