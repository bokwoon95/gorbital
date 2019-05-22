package controllers

import (
	"fmt"
	"net/http"

	"github.com/bokwoon95/orbital/db"
)

// TestDB will ping the database to see if it works
func TestDB(w http.ResponseWriter, r *http.Request) {
	if err := db.DB.Ping(); err != nil {
		fmt.Fprintf(w, err.Error())
		panic(err)
	} else {
		fmt.Fprintf(w, "db connection success")
	}
}
