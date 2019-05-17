package main

import (
	"database/sql"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
)

// TemplateData is an amalgamation of various component data
type TemplateData struct {
	NavbarData *NavbarData
}

// NavbarData contains the necessary information needed to render the navbar
type NavbarData struct {
	LoggedIn              bool
	DisplayName           string
	Role                  string
	ParticipantTeamStatus string
}

func main() {
	r := chi.NewRouter()

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
					Role:                  "public",
					ParticipantTeamStatus: "teamless",
				},
			},
		)
	})

	// "/db"
	r.Get("/db", func(w http.ResponseWriter, r *http.Request) {
		db, err := sql.Open("postgres", "postgres://bokwoon@localhost/orbital_dev?sslmode=disable")
		if err != nil {
			panic(err)
		}
		defer db.Close()
		if err = db.Ping(); err != nil {
			panic(err)
		} else {
			w.Write([]byte("db success"))
		}
	})

	// pastYearShowcase.html "/pys"
	r.Get("/pys", func(w http.ResponseWriter, r *http.Request) {
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
					Role:                  "admin",
					ParticipantTeamStatus: "teamless",
				},
			},
		)
	})

	// responsiveNavbar.html "/nvb"
	r.Get("/nvb", func(w http.ResponseWriter, r *http.Request) {
		mustExecute(
			w,
			mustParse(
				"templates/responsiveNavbar.html",
			),
			nil,
		)
	})

	// hamburger.html "/b"
	r.Get("/b", func(w http.ResponseWriter, r *http.Request) {
		mustExecute(
			w,
			mustParse(
				"templates/burger.html",
			),
			nil,
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
