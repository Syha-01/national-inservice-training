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

// Update updates a specific facilitator.
func (m FacilitatorModel) Update(facilitator *Facilitator) error {
	query := `
		UPDATE facilitators
		SET first_name = $1, last_name = $2, email = $3, personnel_id = $4
		WHERE id = $5
		RETURNING id`

	args := []any{
		facilitator.FirstName,
		facilitator.LastName,
		facilitator.Email,
		sql.NullInt64(facilitator.PersonnelID),
		facilitator.ID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&facilitator.ID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}
	return nil
}
