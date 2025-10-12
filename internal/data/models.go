package data

import "database/sql"

type Models struct {
	Officers OfficerModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Officers: OfficerModel{DB: db},
	}
}
