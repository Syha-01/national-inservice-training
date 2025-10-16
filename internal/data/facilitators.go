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

// Create creates a new facilitator.
func (m FacilitatorModel) Create(facilitator *Facilitator) error {
	query := `
		INSERT INTO facilitators (first_name, last_name, email, personnel_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id`

	args := []any{
		facilitator.FirstName,
		facilitator.LastName,
		facilitator.Email,
		sql.NullInt64(facilitator.PersonnelID),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&facilitator.ID)
}

// GetAll returns a slice of all facilitators.
func (m FacilitatorModel) GetAll() ([]*Facilitator, error) {
	query := `
		SELECT id, first_name, last_name, email, personnel_id
		FROM facilitators
		ORDER BY id`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	facilitators := []*Facilitator{}

	for rows.Next() {
		var facilitator Facilitator
		err := rows.Scan(
			&facilitator.ID,
			&facilitator.FirstName,
			&facilitator.LastName,
			&facilitator.Email,
			(*sql.NullInt64)(&facilitator.PersonnelID),
		)
		if err != nil {
			return nil, err
		}
		facilitators = append(facilitators, &facilitator)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return facilitators, nil
}

// Delete deletes a specific facilitator.
func (m FacilitatorModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
		DELETE FROM facilitators
		WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}

// AssignToSession assigns a facilitator to a session.
func (m FacilitatorModel) AssignToSession(sessionID, facilitatorID int64) error {
	query := `
		INSERT INTO session_facilitators (session_id, facilitator_id)
		VALUES ($1, $2)`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, query, sessionID, facilitatorID)
	return err
}

// RemoveFromSession removes a facilitator from a session.
func (m FacilitatorModel) RemoveFromSession(sessionID, facilitatorID int64) error {
	query := `
		DELETE FROM session_facilitators
		WHERE session_id = $1 AND facilitator_id = $2`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := m.DB.ExecContext(ctx, query, sessionID, facilitatorID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}
