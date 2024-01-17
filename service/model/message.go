package model

import (
	"context"
	"github.com/jmoiron/sqlx"
	"time"
)

type Message struct {
	ID          int            `db:"id" json:"id"`
	SubmittedAt time.Time      `db:"submitted_at" json:"submitted_at"`
	IPAddress   NullableString `db:"ip_address" json:"ip_address"`
	Name        string         `db:"name" json:"name"`
	Email       string         `db:"email" json:"email"`
	Phone       NullableString `db:"phone" json:"phone"`
	OutsideUS   bool           `db:"outside_us" json:"outside_us"`
	Zip         NullableString `db:"zip" json:"zip"`
	City        NullableString `db:"city" json:"city"`
	Message     NullableString `db:"message" json:"message"`
	Status      string         `db:"status" json:"status"`
}

func InsertMessage(ctx context.Context, db *sqlx.DB, message Message) error {
	_, err := db.NamedExecContext(
		ctx,
		`INSERT INTO messages (
                      			ip_address,
                      			name,
                      			email,
                      			phone,
                      			outside_us,
                      			zip,
                      			city,
                      			message,
                      			status	
					  		) VALUES (
								:ip_address,
								:name,
								:email,
								:phone,
								:outside_us,
								:zip,
								:city,
								:message,
								'PENDING'
					  		)`,
		message,
	)
	return err
}
