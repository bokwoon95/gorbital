package controllers

import (
	"net/http"

	auth "github.com/bokwoon95/orbital/auth"
	db "github.com/bokwoon95/orbital/orbital_db"
	erro "github.com/bokwoon95/orbital/erro"
)

// Contract86c89e6 contains the variables that will be passed into register.86c89e6.html
type Contract86c89e6 struct {
	LoggedIn                bool
	DisplayName             string
	Role                    string
	Roles                   db.RolesStruct
	ParticipantTeamStatus   string
	ParticipantTeamStatuses db.ParticipantTeamStatusesStruct
}

// RegisterGet lorem ipsum
func RegisterGet(w http.ResponseWriter, r *http.Request) {
	loggedIn, _, displayName, role, participantTeamStatus, err := db.GetNavbarData(r)
	if err != nil {
		erro.Dump(w, err)
		return
	}

	mustExecute(w, mustParse(w,
		"orbital_views/register.86c89e6.html",
		"orbital_views/navbar.html",
	), &Contract86c89e6{
		LoggedIn:                loggedIn,
		DisplayName:             displayName,
		Role:                    role,
		Roles:                   db.Roles,
		ParticipantTeamStatus:   participantTeamStatus,
		ParticipantTeamStatuses: db.ParticipantTeamStatuses,
	})
}

// RegisterPost lorem ipsum
func RegisterPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		erro.Dump(w, err)
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
		erro.Dump(w, err)
		return
	}

	var uid int
	uid, err = db.InsertParticipant(nusnetid, passwordhash, displayname)
	if err != nil {
		erro.Dump(w, err)
		return
	}

	err = auth.SetSession(w, r, uid)
	if err != nil {
		erro.Dump(w, err)
		return
	}

	http.Redirect(w, r, "/", 301)
}
