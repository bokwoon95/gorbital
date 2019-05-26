package controllers

import (
	"net/http"

	"github.com/bokwoon95/orbital/db"
	"github.com/bokwoon95/orbital/erro"
)

// Contract361d489 contains the variables that will be passed into pastYearShowcase.361d489.html
type Contract361d489 struct {
	LoggedIn                bool
	DisplayName             string
	Role                    string
	Roles                   db.RolesStruct
	ParticipantTeamStatus   string
	ParticipantTeamStatuses db.ParticipantTeamStatusesStruct
	Projects                []db.Project
}

// PastYearShowcase361d489 lorem ipsum
func PastYearShowcase361d489(w http.ResponseWriter, r *http.Request) {
	loggedIn, _, displayName, role, participantTeamStatus, err := db.GetNavbarData(r)
	if err != nil {
		erro.Dump(w, err)
	}

	projects, err := db.GetShowcaseProjects()
	if err != nil {
		erro.Dump(w, err)
	}

	mustExecute(w, mustParse(w,
		"html/pastYearShowcase.361d489.html",
		"html/navbar.html",
	), &Contract361d489{
		LoggedIn:                loggedIn,
		DisplayName:             displayName,
		Role:                    role,
		Roles:                   db.Roles,
		ParticipantTeamStatus:   participantTeamStatus,
		ParticipantTeamStatuses: db.ParticipantTeamStatuses,
		Projects:                projects,
	})
}
