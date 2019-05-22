package controllers

import (
	"fmt"
	"net/http"

	"github.com/bokwoon95/orbital/auth"
	"github.com/davecgh/go-spew/spew"
)

// ContractSession lorem ipsum
type ContractSession struct {
	LoggedIn              bool
	DisplayName           string
	Role                  string
	ParticipantTeamStatus string
	SessionCookie         string
}

// SessionGet will ping the database to see if it works
func SessionGet(w http.ResponseWriter, r *http.Request) {
	hashedCookie, uid := auth.GetActiveSession(r)
	spew.Printf("a:%s b:%d", hashedCookie, uid)
		mustExecute(w, mustParse(
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
	hashedCookie, _ := auth.GetActiveSession(r)
	if hashedCookie == "" {
		err := auth.SetSession(w, r, 1)
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}
	}
	http.Redirect(w, r, "/session", 301)
}

// SessionRevoke lorem ipsum
func SessionRevoke(w http.ResponseWriter, r *http.Request) {
	err := auth.RevokeSession(w, r)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	http.Redirect(w, r, "/session", 301)
	mustExecute(w, mustParse(
		"html/session.html",
		"html/navbar.html",
	), &ContractSession{
		LoggedIn:              false,
		DisplayName:           "",
		Role:                  "",
		ParticipantTeamStatus: "",
	})
}
