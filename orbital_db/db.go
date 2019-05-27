package db

import (
	erro "github.com/bokwoon95/orbital/erro"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // postgres driver for sqlx
)

var (
	// DB represents the database connection object. It is of type *sqlx.DB.
	DB *sqlx.DB
)

// Init will initialize the DB object
func Init() error {
	var err error
	if DB, err = sqlx.Open(
		"postgres",
		"postgres://bokwoon@localhost:5433/orbitaldb_dev?sslmode=disable",
	); err != nil {
		return erro.Wrap(err)
	}
	return nil
}
