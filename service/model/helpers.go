package model

import (
	"database/sql"
	"encoding/json"
)

type NullableString struct {
	sql.NullString
}

func (ns *NullableString) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		ns.Valid = false
		return nil
	}

	ns.Valid = true
	return json.Unmarshal(data, &ns.String)
}
