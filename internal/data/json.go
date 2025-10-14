package data

import (
	"database/sql"
	"encoding/json"
)

type NullInt64 sql.NullInt64

func (ni NullInt64) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Int64)
}
