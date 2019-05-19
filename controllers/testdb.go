package controllers

import (
	"fmt"
	"net/http"

	"github.com/bokwoon95/orbital/models"
)

// TestDB will ping the database to see if it works
func TestDB(w http.ResponseWriter, r *http.Request) {
	if err := models.DB.Ping(); err != nil {
		fmt.Fprintf(w, err.Error())
		panic(err)
	} else {
		fmt.Fprintf(w, "db connection success")
	}
}
