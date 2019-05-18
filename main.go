package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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

func main() {
	// Create router object 'r'
	r := chi.NewRouter()

	// Create database object 'db'
	var db *sqlx.DB
	var err error
	if db, err = sqlx.Open(
		"postgres",
		"postgres://bokwoon@localhost/orbital_dev?sslmode=disable",
	); err != nil {
		fmt.Println("Error when establishing database interface")
		panic(err)
	}
	defer db.Close()

	// aboutOrbital.html "/"
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		mustExecute(
			w,
			mustParse(
				"templates/aboutOrbital.html",
				"templates/navbar.html",
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
	})

	// "/db"
	r.Get("/db", func(w http.ResponseWriter, r *http.Request) {
		if err := db.Ping(); err != nil {
			fmt.Fprintf(w, err.Error())
			panic(err)
		} else {
			fmt.Fprintf(w, "db connection success")
		}
	})

	// pastYearShowcase.html "/showcase"
	r.Get("/showcase", func(w http.ResponseWriter, r *http.Request) {
		rows, _ := db.Queryx(`
		SELECT t.tid ,u.display_name as adviser ,project_name ,ignition_pitch_poster ,project_poster,project_video 
		FROM submissions s JOIN teams t ON t.tid = s.team JOIN advisers a ON a.uid = t.adviser JOIN users u ON u.uid = a.uid;
		`)
		defer rows.Close()
		var projects []Project
		for rows.Next() {
			var project Project
			err = rows.StructScan(&project)
			if project.ProjectPoster.Valid {
				project.SafeProjectPoster = template.URL(project.ProjectPoster.String)
			} else {
				project.SafeProjectPoster = template.URL("")
			}
			if project.ProjectVideo.Valid {
				project.SafeProjectVideo = template.URL(project.ProjectVideo.String)
			} else {
				project.SafeProjectVideo = template.URL("")
			}
			projects = append(projects, project)
		}

		mustExecute(
			w,
			mustParse(
				"templates/pastYearShowcase.html",
				"templates/navbar.html",
			),
			&TemplateData{
				NavbarData: &NavbarData{
					LoggedIn:              true,
					DisplayName:           "User01",
					Role:                  "participant",
					ParticipantTeamStatus: "teamless",
				},
				Projects:  projects,
				DebugStrings: []string{spew.Sdump(projects)},
			},
		)
	})

	// Ensure files in static/ are available to the public
	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "static")
	FileServer(r, "/static", http.Dir(filesDir))

	// Start the server
	http.ListenAndServe(":3000", r)
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

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}
