package controllers

import (
	"net/http"

	"github.com/bokwoon95/orbital/db"
	"github.com/bokwoon95/orbital/erro"
)

// Contract5380c8f contains the variables that will be passed into aboutOrbital.5380c8f.html
type Contract5380c8f struct {
	LoggedIn                bool
	DisplayName             string
	Role                    string
	Roles                   db.RolesStruct
	ParticipantTeamStatus   string
	ParticipantTeamStatuses db.ParticipantTeamStatusesStruct
}

// AboutOrbital5380c8f lorem ipsum
func AboutOrbital5380c8f(w http.ResponseWriter, r *http.Request) {
	loggedIn, _, displayName, role, participantTeamStatus, err := db.GetNavbarData(r)
	if err != nil {
		erro.Dump(w, err)
		return
	}

	mustExecute(w, mustParse(w,
		"html/aboutOrbital.5380c8f.html",
		"html/navbar.html",
	), &Contract5380c8f{
		LoggedIn:                loggedIn,
		DisplayName:             displayName,
		Role:                    role,
		Roles:                   db.Roles,
		ParticipantTeamStatus:   participantTeamStatus,
		ParticipantTeamStatuses: db.ParticipantTeamStatuses,
	})
}
