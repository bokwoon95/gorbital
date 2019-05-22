package controllers

import (
	"database/sql"
	"html/template"
	"net/http"

	"github.com/bokwoon95/orbital/db"
	// "github.com/davecgh/go-spew/spew"
)

// Project lorem ipsum
type Project struct {
	TeamID              string         `db:"tid"`
	Adviser             string         `db:"adviser"`
	ProjectName         string         `db:"project_name"`
	IgnitionPitchPoster string         `db:"ignition_pitch_poster"`
	ProjectPoster       sql.NullString `db:"project_poster"`
	ProjectVideo        sql.NullString `db:"project_video"`
	SafeProjectPoster   template.URL
	SafeProjectVideo    template.URL
}

// Contract361d489 is the contract between PastYearShowcase361d489() and pastYearShowcase.361d489.html
type Contract361d489 struct {
	LoggedIn              bool
	DisplayName           string
	Role                  string
	ParticipantTeamStatus string
	Projects              []Project
	DebugStrings          []string
}

// PastYearShowcase361d489 lorem ipsum
func PastYearShowcase361d489(w http.ResponseWriter, r *http.Request) {
	rows, _ := db.DB.Queryx(`
		SELECT t.tid ,u.display_name as adviser ,project_name ,ignition_pitch_poster ,project_poster,project_video 
		FROM submissions s JOIN teams t ON t.tid = s.team JOIN advisers a ON a.uid = t.adviser JOIN users u ON u.uid = a.uid;
		`)
	defer rows.Close()
	var projects []Project
	for rows.Next() {
		var project Project
		rows.StructScan(&project)
		projects = append(projects, project)
	}

	mustExecute(
		w,
		mustParse(
			"html/pastYearShowcase.361d489.html",
			"html/navbar.html",
		),
		&Contract361d489{
			LoggedIn:              false,
			DisplayName:           "",
			Role:                  "",
			ParticipantTeamStatus: "",
			Projects:              projects,
		},
	)
}
