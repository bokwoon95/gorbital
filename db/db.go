package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // postgres driver for sqlx
)

var (
	// DB represents the database connection object. It is of type *sqlx.DB.
	DB *sqlx.DB
)

// Init will initialize the DB object
func Init() {
	var err error
	if DB, err = sqlx.Open(
		"postgres",
		"postgres://bokwoon@localhost:5433/orbitaldb_dev?sslmode=disable",
	); err != nil {
		panic(err)
	}
}
