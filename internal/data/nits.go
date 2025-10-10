package data

import (
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

func ValidateNit(v *validator.Validator, nit *Nit) {
	v.Check(nit.CourseID > 0, "course_id", "must be provided and be a positive integer")
	v.Check(!nit.StartDate.IsZero(), "start_date", "must be provided")
	v.Check(!nit.EndDate.IsZero(), "end_date", "must be provided")
	v.Check(nit.EndDate.After(nit.StartDate), "end_date", "must be after start date")
	v.Check(nit.Location != "", "location", "must be provided")
	v.Check(len(nit.Location) <= 100, "location", "must not be more than 100 bytes long")
}
