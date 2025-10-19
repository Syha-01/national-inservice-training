package data

import (
	"database/sql"
)

type Models struct {
	Officers     OfficerModel
	Courses      CourseModel
	Facilitators FacilitatorModel
	Feedback     FeedbackModel
	Nits         NitModel
	Users        UserModel
	Tokens       TokenModel
	Permissions PermissionModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Officers:     OfficerModel{DB: db},
		Courses:      CourseModel{DB: db},
		Facilitators: FacilitatorModel{DB: db},
		Feedback:     FeedbackModel{DB: db},
		Nits:         NitModel{DB: db},
		Users:        UserModel{DB: db},
		Tokens:       TokenModel{DB: db},
		Permissions: PermissionModel{DB: db},
	}
}
