package data

import (
	"time"
)

// Nit defines the structure for a national inservice training session.
type Nit struct {
	ID        int64     `json:"id"`
	CourseID  int64     `json:"course_id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Location  string    `json:"location"`
	CreatedAt time.Time `json:"created_at"`
	Version   int32     `json:"version"`
}
