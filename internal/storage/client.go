package storage

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Client struct {
	conn *sqlx.DB
}

func NewClient(db *sql.DB) (*Client, error) {
	sqlxDB := sqlx.NewDb(db, "postgres")

	errPing := db.Ping()
	if errPing != nil {
		return nil, errPing
	}

	return &Client{
			conn: sqlxDB,
		},
		nil
}
