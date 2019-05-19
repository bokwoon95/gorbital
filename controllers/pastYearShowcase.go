package controllers

import (
	"net/http"

	"github.com/bokwoon95/orbital/models"
	"github.com/davecgh/go-spew/spew"
)

// PastYearShowcase haha
func PastYearShowcase(w http.ResponseWriter, r *http.Request) {
	rows, _ := models.DB.Queryx(`
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
			"html/pastYearShowcase.html",
			"html/navbar.html",
		),
		&TemplateData{
			NavbarData: &NavbarData{
				LoggedIn:              true,
				DisplayName:           "User01",
				Role:                  "participant",
				ParticipantTeamStatus: "teamless",
			},
			Projects:     projects,
			DebugStrings: []string{spew.Sdump(projects)},
		},
	)
}
