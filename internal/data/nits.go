package data

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/Syha-01/national-inservice-training/internal/validator"
)

// Nit defines the structure for a national inservice training session.
type Nit struct {
	ID        int64     `json:"id"`
	CourseID  int64     `json:"course_id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Location  string    `json:"location"`
	CreatedAt time.Time `json:"-"`
	Version   int32     `json:"version"`
}

// Officer defines the structure for a police officer.
type Officer struct {
	ID               int64     `json:"id"`
	RegulationNumber string    `json:"regulation_number"`
	FirstName        string    `json:"first_name"`
	LastName         string    `json:"last_name"`
	Sex              string    `json:"sex"`
	RankID           int64     `json:"rank_id,omitempty"`
	FormationID      int64     `json:"formation_id,omitempty"`
	PostingID        int64     `json:"posting_id,omitempty"`
	IsActive         bool      `json:"is_active"`
	CreatedAt        time.Time `json:"-"`
	UpdatedAt        time.Time `json:"-"`
}

func ValidateNit(v *validator.Validator, nit *Nit) {
	v.Check(nit.CourseID > 0, "course_id", "must be provided and be a positive integer")
	v.Check(!nit.StartDate.IsZero(), "start_date", "must be provided")
	v.Check(!nit.EndDate.IsZero(), "end_date", "must be provided")
	v.Check(nit.EndDate.After(nit.StartDate), "end_date", "must be after start date")
	v.Check(nit.Location != "", "location", "must be provided")
	v.Check(len(nit.Location) <= 100, "location", "must not be more than 100 bytes long")
}

func ValidateOfficer(v *validator.Validator, officer *Officer) {
	v.Check(officer.RegulationNumber != "", "regulation_number", "must be provided")
	v.Check(len(officer.RegulationNumber) <= 50, "regulation_number", "must not be more than 50 bytes long")
	v.Check(officer.FirstName != "", "first_name", "must be provided")
	v.Check(len(officer.FirstName) <= 100, "first_name", "must not be more than 100 bytes long")
	v.Check(officer.LastName != "", "last_name", "must be provided")
	v.Check(len(officer.LastName) <= 100, "last_name", "must not be more than 100 bytes long")
	v.Check(officer.Sex == "Male" || officer.Sex == "Female", "sex", "must be either Male or Female")
}

// OfficerModel wraps the database connection pool.
type OfficerModel struct {
	DB *sql.DB
}

// Get a specific Officer from the personnel table
func (m OfficerModel) Get(id int64) (*Officer, error) {
	// check if the id is valid
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	// the SQL query to be executed against the database table
	query := `
		SELECT id, regulation_number, first_name, last_name, sex, rank_id, formation_id, posting_id, is_active, created_at, updated_at
		FROM personnel
		WHERE id = $1
`
	// declare a variable of type Officer to store the returned officer
	var officer Officer
	// Set a 3-second context/timer
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&officer.ID,
		&officer.RegulationNumber,
		&officer.FirstName,
		&officer.LastName,
		&officer.Sex,
		&officer.RankID,
		&officer.FormationID,
		&officer.PostingID,
		&officer.IsActive,
		&officer.CreatedAt,
		&officer.UpdatedAt,
	)
	// check for which type of error
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &officer, nil
}

// Update a specific Officer in the personnel table
func (m OfficerModel) Update(officer *Officer) error {
	query := `
		UPDATE personnel
		SET regulation_number = $1, first_name = $2, last_name = $3, sex = $4, rank_id = $5, formation_id = $6, posting_id = $7, is_active = $8, updated_at = NOW()
		WHERE id = $9
`
	args := []any{
		officer.RegulationNumber,
		officer.FirstName,
		officer.LastName,
		officer.Sex,
		officer.RankID,
		officer.FormationID,
		officer.PostingID,
		officer.IsActive,
		officer.ID,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, query, args...)
	return err
}
