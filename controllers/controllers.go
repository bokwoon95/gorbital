package controllers

import (
	"html/template"
	"net/http"

	"github.com/bokwoon95/orbital/erro"
)

func mustParse(w http.ResponseWriter, templateFiles ...string) *template.Template {
	t, err := template.ParseFiles(templateFiles...)
	if err != nil {
		erro.Dump(w, err)
	}
	return t
}

func mustExecute(w http.ResponseWriter, t *template.Template, data interface{}) {
	if err := t.Execute(w, data); err != nil {
		erro.Dump(w, err)
	}
}
