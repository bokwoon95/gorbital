package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bokwoon95/orbital/auth"
	"github.com/bokwoon95/orbital/db"
	"github.com/davecgh/go-spew/spew"
)

// Contract689a9b4 lorem ipsum
type Contract689a9b4 struct {
	LoggedIn              bool
	DisplayName           string
	Role                  string
	ParticipantTeamStatus string
	DebugString           string
}

// LoginGet689a9b4 lorem ipsum
func LoginGet689a9b4(w http.ResponseWriter, r *http.Request) {
	mustExecute(w,
		mustParse(
			"html/login.689a9b4.html",
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

// LoginPost689a9b4 lorem ipsum
func LoginPost689a9b4(w http.ResponseWriter, r *http.Request) {
	var err error

	err = r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	nusnetid := r.FormValue("nusnetid")
	password := r.FormValue("password")
	type temp struct {
		UserID       string `db:"uid"`
		PasswordHash string `db:"password"`
		DisplayName  string `db:"display_name"`
	}
	var t temp
	err = db.DB.QueryRowx("SELECT uid, display_name, password FROM users WHERE nusnetid = $1", nusnetid).StructScan(&t)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, err.Error())
		return
	}

	err = auth.CompareHashAndPassword(t.PasswordHash, password)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	mustExecute(w, mustParse(
			"html/login.689a9b4.html",
			"html/navbar.html",
		), &Contract86c89e6{
			LoggedIn:              false,
			DisplayName:           "",
			Role:                  "public",
			ParticipantTeamStatus: "teamless",
			DebugString:           "congrats, you're in " + t.DisplayName,
		},
	)
}
