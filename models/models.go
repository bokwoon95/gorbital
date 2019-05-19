package models

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // postgres driver for sqlx
)

var (
	// DB represents the database connection object
	DB *sqlx.DB
)

// InitDB will initialize the DB object
func InitDB() {
	var err error
	if DB, err = sqlx.Open(
		"postgres",
		"postgres://bokwoon@localhost/orbital_dev?sslmode=disable",
	); err != nil {
		fmt.Println("Error when establishing database interface")
		panic(err)
	}
}
