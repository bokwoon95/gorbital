package db

import (
	"database/sql"

	"github.com/bokwoon95/orbital/erro"
)

// Project lorem ipsum
type Project struct {
	Milestone           string         `db:"milestone"`
	ProjectLevel        string         `db:"project_level"`
	ProjectName         string         `db:"project_name"`
	ProjectLink         sql.NullString `db:"project_link"`
	ProjectReadme       sql.NullString `db:"project_readme"`
	ProjectPoster       sql.NullString `db:"project_poster"`
	ProjectVideo        sql.NullString `db:"project_video"`
	Cohort              string         `db:"cohort"`
	TeamProjectLevel    string         `db:"team_project_level"`
	TeamName            string         `db:"team_name"`
	Participant1        string         `db:"participant1"`
	Participant2        string         `db:"participant2"`
	Adviser             sql.NullString `db:"adviser"`
	Mentor              sql.NullString `db:"mentor"`
	IgnitionPitchPoster sql.NullString `db:"ignition_pitch_poster"`
}

// GetShowcaseProjects lorem ipsum
func GetShowcaseProjects() ([]Project, error) {
	var projects []Project
	rows, err := DB.Queryx(`
SELECT
	milestone
	,project_level
	,project_name
	,project_link
	,project_readme
	,project_poster
	,project_video
	,cohort
	,team_project_level
	,team_name
	,participant1
	,participant2
	,adviser
	,mentor
	,ignition_pitch_poster
FROM
	v_submissions
	`)
	if err != nil {
		erro.WrapX(&err)
	}
	defer rows.Close()
	for rows.Next() {
		var project Project
		err = rows.StructScan(&project)
		if err != nil {
			erro.WrapX(&err)
			break
		}
		projects = append(projects, project)
	}
	return projects, err
}
