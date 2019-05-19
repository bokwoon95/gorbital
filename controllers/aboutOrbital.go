package controllers

import (
	"net/http"
)

// AboutOrbital haha
func AboutOrbital(w http.ResponseWriter, r *http.Request) {
	mustExecute(
		w,
		mustParse(
			"html/aboutOrbital.html",
			"html/navbar.html",
		),
		&TemplateData{
			NavbarData: &NavbarData{
				LoggedIn:              true,
				DisplayName:           "User01",
				Role:                  "participant",
				ParticipantTeamStatus: "teamless",
			},
		},
	)
}
