package main

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/bokwoon95/orbital/controllers"
	"github.com/bokwoon95/orbital/models"
	"github.com/go-chi/chi"
)

func main() {
	// Initialize router
	r := chi.NewRouter()

	// Initialize database
	models.InitDB()
	defer models.DB.Close() // When the main function dies, so should the database

	// Set up routes
	r.Get("/", controllers.AboutOrbital)
	r.Get("/testdb", controllers.TestDB)
	r.Get("/showcase", controllers.PastYearShowcase)

	// Ensure these directories are accessible to the public
	workDir, _ := os.Getwd()
	FileServer(r, "/images", http.Dir(filepath.Join(workDir, "/images")))
	FileServer(r, "/css", http.Dir(filepath.Join(workDir, "/css")))
	FileServer(r, "/js", http.Dir(filepath.Join(workDir, "/js")))

	// Start the server
	http.ListenAndServe(":3000", r)
}

// FileServer conveniently sets up a http.FileServer handler to serve static files from a http.FileSystem.
// See https://github.com/go-chi/chi/blob/18d990c0d1c023b05a3652d322ae36d8bdb62e07/_examples/fileserver/main.go
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
