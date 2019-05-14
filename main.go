package main

import (
	"html/template"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	// About Orbital
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		mustExecute(
			w,
			mustParse(
				"templates/aboutOrbital.html",
				"templates/navbars.html",
			),
			nil,
		)
	})

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
