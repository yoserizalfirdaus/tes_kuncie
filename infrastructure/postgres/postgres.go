package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect(datasource string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", datasource)
	if err != nil {
		return nil, err
	}

	return db, nil
}
