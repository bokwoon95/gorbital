package controllers

import (
	"net/http"

	"github.com/bokwoon95/orbital/auth"
	"github.com/bokwoon95/orbital/db"
	"github.com/bokwoon95/orbital/erro"
)

// ContractSession lorem ipsum
type ContractSession struct {
	LoggedIn                bool
	DisplayName             string
	Role                    string
	Roles                   db.RolesStruct
	ParticipantTeamStatus   string
	ParticipantTeamStatuses db.ParticipantTeamStatusesStruct
	SessionCookie           string
}

// SessionGet will ping the database to see if it works
func SessionGet(w http.ResponseWriter, r *http.Request) {
	hashedCookie, uid, _ := auth.GetActiveSession(r)
	mustExecute(w, mustParse(w,
		"html/session.html",
		"html/navbar.html",
	), &ContractSession{
		LoggedIn:              hashedCookie != "",
		DisplayName:           string(uid),
		Role:                  "",
		ParticipantTeamStatus: "",
	})
}

// SessionSet lorem ipsum
func SessionSet(w http.ResponseWriter, r *http.Request) {
	hashedCookie, _, _ := auth.GetActiveSession(r)
	if hashedCookie == "" {
		err := auth.SetSession(w, r, 1)
		if err != nil {
			erro.Dump(w, err)
		}
	}
	http.Redirect(w, r, "/session", 301)
}

// SessionRevoke lorem ipsum
func SessionRevoke(w http.ResponseWriter, r *http.Request) {
	err := auth.RevokeSession(w, r)
	if err != nil {
		erro.Dump(w, err)
		return
	}
	http.Redirect(w, r, "/", 301)
}
