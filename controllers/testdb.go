package controllers

import (
	"net/http"

	"github.com/bokwoon95/orbital/db"
	"github.com/bokwoon95/orbital/erro"
)

// TestDB will ping the database to see if it works
func TestDB(w http.ResponseWriter, r *http.Request) {
	if err := db.DB.Ping(); err != nil {
		erro.Dump(w, err)
	} else {
		w.Write([]byte("db connection success"))
	}
}
