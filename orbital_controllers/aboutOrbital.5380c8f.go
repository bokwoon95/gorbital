package controllers

import (
	"net/http"

	db "github.com/bokwoon95/orbital/orbital_db"
	erro "github.com/bokwoon95/orbital/erro"
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

// AboutOrbital lorem ipsum
func AboutOrbital(w http.ResponseWriter, r *http.Request) {
	loggedIn, _, displayName, role, participantTeamStatus, err := db.GetNavbarData(r)
	if err != nil {
		erro.Dump(w, err)
		return
	}

	mustExecute(w, mustParse(w,
		"orbital_views/aboutOrbital.5380c8f.html",
		"orbital_views/navbar.html",
	), &Contract5380c8f{
		LoggedIn:                loggedIn,
		DisplayName:             displayName,
		Role:                    role,
		Roles:                   db.Roles,
		ParticipantTeamStatus:   participantTeamStatus,
		ParticipantTeamStatuses: db.ParticipantTeamStatuses,
	})
}
