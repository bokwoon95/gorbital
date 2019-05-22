package controllers

import (
	"net/http"
	// "github.com/davecgh/go-spew/spew"
)

// Contract5380c8f is the contract between AboutOrbital5380c8f() and aboutOrbital.5380c8f.html
type Contract5380c8f struct {
	LoggedIn              bool
	DisplayName           string
	Role                  string
	ParticipantTeamStatus string
}

// AboutOrbital5380c8f lorem ipsum
func AboutOrbital5380c8f(w http.ResponseWriter, r *http.Request) {
	mustExecute(
		w,
		mustParse(
			"html/aboutOrbital.5380c8f.html",
			"html/navbar.html",
		),
		&Contract5380c8f{
			LoggedIn:              false,
			DisplayName:           "",
			Role:                  "public",
			ParticipantTeamStatus: "",
		},
	)
}
