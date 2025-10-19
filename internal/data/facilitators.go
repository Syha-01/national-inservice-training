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
	Version     int32     `json:"version"`
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
		SELECT id, first_name, last_name, email, personnel_id, version
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
		&facilitator.Version,
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

// GetByPersonnelID retrieves a specific facilitator by personnel ID.
func (m FacilitatorModel) GetByPersonnelID(id int64) (*Facilitator, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
		SELECT id, first_name, last_name, email, personnel_id, version
		FROM facilitators
		WHERE personnel_id = $1`

	var facilitator Facilitator

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&facilitator.ID,
		&facilitator.FirstName,
		&facilitator.LastName,
		&facilitator.Email,
		(*sql.NullInt64)(&facilitator.PersonnelID),
		&facilitator.Version,
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
		SET first_name = $1, last_name = $2, email = $3, personnel_id = $4, version = version + 1
		WHERE id = $5 AND version = $6
		RETURNING version`

	args := []any{
		facilitator.FirstName,
		facilitator.LastName,
		facilitator.Email,
		sql.NullInt64(facilitator.PersonnelID),
		facilitator.ID,
		facilitator.Version,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&facilitator.Version)
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
		RETURNING id, version`

	args := []any{
		facilitator.FirstName,
		facilitator.LastName,
		facilitator.Email,
		sql.NullInt64(facilitator.PersonnelID),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&facilitator.ID, &facilitator.Version)
}

// GetAll returns a slice of all facilitators.
func (m FacilitatorModel) GetAll(filters Filters) ([]*Facilitator, Metadata, error) {
	query := `
		SELECT COUNT(*) OVER(), id, first_name, last_name, email, personnel_id, version
		FROM facilitators
		ORDER BY id
		LIMIT $1 OFFSET $2`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, filters.limit(), filters.offset())
	if err != nil {
		return nil, Metadata{}, err
	}
	defer rows.Close()

	totalRecords := 0
	facilitators := []*Facilitator{}

	for rows.Next() {
		var facilitator Facilitator
		err := rows.Scan(
			&totalRecords,
			&facilitator.ID,
			&facilitator.FirstName,
			&facilitator.LastName,
			&facilitator.Email,
			(*sql.NullInt64)(&facilitator.PersonnelID),
			&facilitator.Version,
		)
		if err != nil {
			return nil, Metadata{}, err
		}
		facilitators = append(facilitators, &facilitator)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return facilitators, metadata, nil
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

// GetAllForSession returns a slice of all facilitators for a specific session.
func (m FacilitatorModel) GetAllForSession(sessionID int64, filters Filters) ([]*Facilitator, Metadata, error) {
	query := `
		SELECT COUNT(*) OVER(), f.id, f.first_name, f.last_name, f.email, f.personnel_id, f.version
		FROM facilitators f
		INNER JOIN session_facilitators sf ON f.id = sf.facilitator_id
		WHERE sf.session_id = $1
		ORDER BY f.id
		LIMIT $2 OFFSET $3`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, sessionID, filters.limit(), filters.offset())
	if err != nil {
		return nil, Metadata{}, err
	}
	defer rows.Close()

	totalRecords := 0
	facilitators := []*Facilitator{}

	for rows.Next() {
		var facilitator Facilitator
		err := rows.Scan(
			&totalRecords,
			&facilitator.ID,
			&facilitator.FirstName,
			&facilitator.LastName,
			&facilitator.Email,
			(*sql.NullInt64)(&facilitator.PersonnelID),
			&facilitator.Version,
		)
		if err != nil {
			return nil, Metadata{}, err
		}
		facilitators = append(facilitators, &facilitator)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return facilitators, metadata, nil
}

// AssignToSession assigns a facilitator to a session.
func (m FacilitatorModel) AssignToSession(sessionID, facilitatorID int64) error {
	query := `
		INSERT INTO session_facilitators (session_id, facilitator_id)
		SELECT $1, $2
		WHERE NOT EXISTS (
			SELECT 1
			FROM session_facilitators
			WHERE session_id = $1 AND facilitator_id = $2
		)`

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
		return errors.New("facilitator already assigned to this session")
	}

	return nil
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
