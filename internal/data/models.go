package data

import (
	"database/sql"
)

type Models struct {
	Officers     OfficerModel
	Courses      CourseModel
	Facilitators FacilitatorModel
	Feedback     FeedbackModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Officers:     OfficerModel{DB: db},
		Courses:      CourseModel{DB: db},
		Facilitators: FacilitatorModel{DB: db},
		Feedback:     FeedbackModel{DB: db},
	}
}
