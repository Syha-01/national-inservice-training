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

// Course defines the structure for a training course.
type Course struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Category    string    `json:"category"` // 'Mandatory' or 'Elective'
	CreditHours float64   `json:"credit_hours"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
	Version     int32     `json:"version"`
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

// ValidateCourse validates a Course struct
func ValidateCourse(v *validator.Validator, course *Course) {
	v.Check(course.Title != "", "title", "must be provided")
	v.Check(len(course.Title) <= 255, "title", "must not be more than 255 bytes long")
	v.Check(course.Category != "", "category", "must be provided")
	v.Check(course.Category == "Mandatory" || course.Category == "Elective", "category", "must be either 'Mandatory' or 'Elective'")
	v.Check(course.CreditHours > 0, "credit_hours", "must be greater than 0")
	v.Check(len(course.Description) <= 1000, "description", "must not be more than 1000 bytes long")
}

// OfficerModel wraps the database connection pool.
type OfficerModel struct {
	DB *sql.DB
}

// CourseModel wraps the database connection pool.
type CourseModel struct {
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

// Delete a specific Officer from the personnel table
func (m OfficerModel) Delete(id int64) error {
	// check if the id is valid
	if id < 1 {
		return ErrRecordNotFound
	}
	// the SQL query to be executed against the database table
	query := `
		DELETE FROM personnel
		WHERE id = $1
		`
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

// GetAll retrieves all courses from the database
func (m CourseModel) GetAll() ([]*Course, error) {
	query := `
		SELECT id, title, description, category, credit_hours, created_at, updated_at
		FROM courses
		ORDER BY id
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	courses := []*Course{}

	for rows.Next() {
		var course Course
		err := rows.Scan(
			&course.ID,
			&course.Title,
			&course.Description,
			&course.Category,
			&course.CreditHours,
			&course.CreatedAt,
			&course.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		courses = append(courses, &course)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return courses, nil
}

// Get retrieves a specific course by ID
func (m CourseModel) Get(id int64) (*Course, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
		SELECT id, title, description, category, credit_hours, created_at, updated_at
		FROM courses
		WHERE id = $1
	`

	var course Course

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&course.ID,
		&course.Title,
		&course.Description,
		&course.Category,
		&course.CreditHours,
		&course.CreatedAt,
		&course.UpdatedAt,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &course, nil
}

// Create inserts a new course into the database
func (m CourseModel) Create(course *Course) error {
	query := `
		INSERT INTO courses (title, description, category, credit_hours, created_at, updated_at)
		VALUES ($1, $2, $3, $4, NOW(), NOW())
		RETURNING id, created_at, updated_at
	`

	args := []any{
		course.Title,
		course.Description,
		course.Category,
		course.CreditHours,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(
		&course.ID,
		&course.CreatedAt,
		&course.UpdatedAt,
	)

	return err
}

// Update modifies an existing course in the database
func (m CourseModel) Update(course *Course) error {
	query := `
		UPDATE courses
		SET title = $1, description = $2, category = $3, credit_hours = $4, updated_at = NOW()
		WHERE id = $5
		RETURNING updated_at
	`

	args := []any{
		course.Title,
		course.Description,
		course.Category,
		course.CreditHours,
		course.ID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&course.UpdatedAt)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrRecordNotFound
		default:
			return err
		}
	}

	return nil
}

// Delete removes a course from the database
func (m CourseModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
		DELETE FROM courses
		WHERE id = $1
	`

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
