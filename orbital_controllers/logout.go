package controllers

import (
	"net/http"

	auth "github.com/bokwoon95/orbital/auth"
	erro "github.com/bokwoon95/orbital/erro"
)

// Logout lorem ipsum
func Logout(w http.ResponseWriter, r *http.Request) {
	err := auth.RevokeSession(w, r)
	if err != nil {
		erro.Dump(w, err)
		return
	}
	http.Redirect(w, r, "/", 301)
}
