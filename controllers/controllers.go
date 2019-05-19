package controllers

import (
	"database/sql"
	"html/template"
	"net/http"
)

// TemplateData is an amalgamation of various component data
type TemplateData struct {
	NavbarData   *NavbarData
	Projects     []Project
	DebugStrings []string
}

// NavbarData contains the necessary information needed to render the navbar
type NavbarData struct {
	LoggedIn              bool
	DisplayName           string `db:"display_name"`
	Role                  string
	ParticipantTeamStatus string
}

// Project contains the necessary information needed to render each project showcase
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

func mustParse(templateFiles ...string) *template.Template {
	t, err := template.ParseFiles(templateFiles...)
	if err != nil {
		panic(err)
	}
	return t
}

func mustExecute(w http.ResponseWriter, t *template.Template, data interface{}) {
	if err := t.Execute(w, data); err != nil {
		panic(err)
	}
}
