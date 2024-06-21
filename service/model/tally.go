package model

import "github.com/jmoiron/sqlx"

func GetTally(db *sqlx.DB, campaign string) (int, error) {
	var tally int
	err := db.Get(
		&tally,
		`SELECT count(*) as total
		FROM messages
		WHERE campaign = ?`,
		campaign,
	)
	if (err != nil) {
		return 0, err
	}
	return tally, nil 
}