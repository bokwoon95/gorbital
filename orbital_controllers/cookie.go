package controllers

import "net/http"

// CookieInspectorGet lorem ipsum
func CookieInspectorGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!"))
}

// CookieInspectorPost lorem ipsum
func CookieInspectorPost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!"))
}
