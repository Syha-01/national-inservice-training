package data

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

// Facilitator defines the structure for a facilitator.
type Facilitator struct {
	ID          int64     `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	PersonnelID NullInt64 `json:"personnel_id,omitempty"`
}

// FacilitatorModel wraps the database connection pool.
type FacilitatorModel struct {
	DB *sql.DB
}

// Get retrieves a specific facilitator by ID.
func (m FacilitatorModel) Get(id int64) (*Facilitator, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
		SELECT id, first_name, last_name, email, personnel_id
		FROM facilitators
		WHERE id = $1`

	var facilitator Facilitator

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&facilitator.ID,
		&facilitator.FirstName,
		&facilitator.LastName,
		&facilitator.Email,
		(*sql.NullInt64)(&facilitator.PersonnelID),
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &facilitator, nil
}
