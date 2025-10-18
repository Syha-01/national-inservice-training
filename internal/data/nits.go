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

// NitModel wraps the database connection pool.
type NitModel struct {
	DB *sql.DB
}

// Get retrieves a specific training session by ID.
func (m NitModel) Get(id int64) (*Nit, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
		SELECT id, course_id, start_date, end_date, location, created_at, version
		FROM training_sessions
		WHERE id = $1`

	var nit Nit

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&nit.ID,
		&nit.CourseID,
		&nit.StartDate,
		&nit.EndDate,
		&nit.Location,
		&nit.CreatedAt,
		&nit.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &nit, nil
}

// GetOfficer retrieves a specific officer by ID.
func (m OfficerModel) GetOfficer(id int64) (*Officer, error) {
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

// UpdateOfficer updates a specific officer's details.
func (m OfficerModel) UpdateOfficer(officer *Officer) error {
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

// GetAllOfficers retrieves all officers from the personnel table.
func (m OfficerModel) GetAllOfficers(filters Filters) ([]*Officer, Metadata, error) {
	query := `
		SELECT COUNT(*) OVER(), id, regulation_number, first_name, last_name, sex, rank_id, formation_id, posting_id, is_active, created_at, updated_at
		FROM personnel
		ORDER BY id
		LIMIT $1 OFFSET $2
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, filters.limit(), filters.offset())
	if err != nil {
		return nil, Metadata{}, err
	}
	defer rows.Close()

	totalRecords := 0
	officers := []*Officer{}

	for rows.Next() {
		var officer Officer
		err := rows.Scan(
			&totalRecords,
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
		if err != nil {
			return nil, Metadata{}, err
		}
		officers = append(officers, &officer)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return officers, metadata, nil
}

// DeleteOfficer removes a specific officer from the database.
func (m OfficerModel) DeleteOfficer(id int64) error {
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

// CreateOfficer creates a new officer in the database.
func (m OfficerModel) CreateOfficer(officer *Officer) error {
	query := `
		INSERT INTO personnel (regulation_number, first_name, last_name, sex, rank_id, formation_id, posting_id, is_active)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, created_at, updated_at
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
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&officer.ID, &officer.CreatedAt, &officer.UpdatedAt)
}

// GetAllCourses retrieves all courses from the database.
func (m CourseModel) GetAllCourses(filters Filters) ([]*Course, Metadata, error) {
	query := `
		SELECT COUNT(*) OVER(), id, title, description, category, credit_hours, created_at, updated_at
		FROM courses
		ORDER BY id
		LIMIT $1 OFFSET $2
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, filters.limit(), filters.offset())
	if err != nil {
		return nil, Metadata{}, err
	}
	defer rows.Close()

	totalRecords := 0
	courses := []*Course{}

	for rows.Next() {
		var course Course
		err := rows.Scan(
			&totalRecords,
			&course.ID,
			&course.Title,
			&course.Description,
			&course.Category,
			&course.CreditHours,
			&course.CreatedAt,
			&course.UpdatedAt,
		)
		if err != nil {
			return nil, Metadata{}, err
		}
		courses = append(courses, &course)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return courses, metadata, nil
}

// GetCourse retrieves a specific course by ID.
func (m CourseModel) GetCourse(id int64) (*Course, error) {
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

// CreateCourse inserts a new course into the database.
func (m CourseModel) CreateCourse(course *Course) error {
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

// UpdateCourse modifies an existing course in the database.
func (m CourseModel) UpdateCourse(course *Course) error {
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

// DeleteCourse removes a course from the database.
func (m CourseModel) DeleteCourse(id int64) error {
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
