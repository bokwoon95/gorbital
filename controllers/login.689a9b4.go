package controllers

import (
	"net/http"

	"github.com/bokwoon95/orbital/auth"
	"github.com/bokwoon95/orbital/db"
	"github.com/bokwoon95/orbital/erro"
)

// Contract689a9b4 contains the variables that will be passed into login.689a9b4.html
type Contract689a9b4 struct {
	LoggedIn                bool
	DisplayName             string
	Role                    string
	Roles                   db.RolesStruct
	ParticipantTeamStatus   string
	ParticipantTeamStatuses db.ParticipantTeamStatusesStruct
}

// LoginGet689a9b4 lorem ipsum
func LoginGet689a9b4(w http.ResponseWriter, r *http.Request) {
	// TODO: if loggedIn, redirect to homepage

	loggedIn, _, displayName, role, participantTeamStatus, err := db.GetNavbarData(r)
	if err != nil {
		erro.Dump(w, err)
	}

	mustExecute(w, mustParse(w,
		"html/login.689a9b4.html",
		"html/navbar.html",
	), &Contract689a9b4{
		LoggedIn:                loggedIn,
		DisplayName:             displayName,
		Role:                    role,
		Roles:                   db.Roles,
		ParticipantTeamStatus:   participantTeamStatus,
		ParticipantTeamStatuses: db.ParticipantTeamStatuses,
	})
}

// LoginPost689a9b4 lorem ipsum
func LoginPost689a9b4(w http.ResponseWriter, r *http.Request) {
	// TODO: if loggedIn, ignore request

	err := r.ParseForm()
	if err != nil {
		erro.Dump(w, err)
		return
	}

	nusnetid := r.FormValue("nusnetid")
	password := r.FormValue("password")
	if nusnetid == "" {
		erro.Dump(w, err)
		return
	}
	user, err := db.GetUserByNUSNET(nusnetid)
	if err != nil {
		erro.Dump(w, err)
		return
	}

	err = auth.CompareHashAndPassword(user.PasswordHash, password)
	if err != nil {
		erro.Dump(w, err)
		return
	}

	err = auth.SetSession(w, r, user.ID)
	if err != nil {
		erro.Dump(w, err)
		return
	}

	http.Redirect(w, r, "/", 301)
}
