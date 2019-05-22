package controllers

import (
	"fmt"
	"net/http"

	"github.com/bokwoon95/orbital/auth"
	"github.com/bokwoon95/orbital/db"
	"github.com/davecgh/go-spew/spew"
)

// Contract86c89e6 lorem ipsum
type Contract86c89e6 struct {
	LoggedIn              bool
	DisplayName           string
	Role                  string
	ParticipantTeamStatus string
	DebugString           string
}

// RegisterGet86c89e6 lorem ipsum
func RegisterGet86c89e6(w http.ResponseWriter, r *http.Request) {
	mustExecute(w,
		mustParse(
			"html/register.86c89e6.html",
			"html/navbar.html",
		),
		&Contract86c89e6{
			LoggedIn:              false,
			DisplayName:           "",
			Role:                  "public",
			ParticipantTeamStatus: "teamless",
			DebugString:           spew.Sdump(r),
		},
	)
}

// RegisterPost86c89e6 lorem ipsum
func RegisterPost86c89e6(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	// Accessing a nonexistent key will panic with an unhelpful message
	// TODO: wrap accessing the form values in a .Get() function that will
	// throw helpful errors when trying to access nonexistent keys
	nusnetid := r.FormValue("nusnetid")
	password := r.FormValue("password")
	displayname := r.Form["display_name"][0]
	passwordhash, err := auth.HashPassword(password)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	_, err = db.DB.Exec(`
	INSERT INTO users (nusnetid, password, display_name) VALUES ($1, $2, $3)
	`, nusnetid, passwordhash, displayname)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	mustExecute(w, mustParse(
		"html/register.86c89e6.html",
		"html/navbar.html",
	), &Contract86c89e6{
		LoggedIn:              false,
		DisplayName:           "",
		Role:                  "public",
		ParticipantTeamStatus: "teamless",
		DebugString:           spew.Sdump(r),
	})
}
