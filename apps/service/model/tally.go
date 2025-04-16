package model

import (
	"github.com/jmoiron/sqlx"
)

func GetTally(db *sqlx.DB, campaign string, zipCodes []string) (int, error) {
	var tally int
	query := `SELECT count(*) FROM messages WHERE campaign = ?`
	args := []interface{}{campaign}

	if len(zipCodes) > 0 {
		query += " AND zip IN (?)"
		var err error
		query, args, err = sqlx.In(query, append(args, zipCodes)...)
		if err != nil {
			return 0, err
		}
		query = db.Rebind(query)
	}

	err := db.Get(&tally, query, args...)
	if err != nil {
		return 0, err
	}
	return tally, nil
}
