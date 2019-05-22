package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/bokwoon95/orbital/controllers"
	"github.com/bokwoon95/orbital/db"
	"github.com/bokwoon95/orbital/auth"
	"github.com/go-chi/chi"
)

func main() {
	// Initialize database and defer its closing to the end of main()
	db.Init()
	defer db.DB.Close()

	// Initialize session database
	auth.Init()
	defer auth.SessionDB.Close()


	// configure log to print '<date> <time> <filename:linenumber> message'
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Initialize router
	r := chi.NewRouter()

	// Set up routes
	r.Get("/testdb", controllers.TestDB)
	r.Get("/", controllers.AboutOrbital5380c8f)
	r.Get("/showcase", controllers.PastYearShowcase361d489)
	r.Get("/register", controllers.RegisterGet86c89e6)
	r.Post("/register", controllers.RegisterPost86c89e6)
	r.Get("/login", controllers.LoginGet689a9b4)
	r.Post("/login", controllers.LoginPost689a9b4)

	r.Get("/session", controllers.SessionGet)
	r.Post("/sessionset", controllers.SessionSet)
	r.Post("/sessionrevoke", controllers.SessionRevoke)

	// Ensure the below directories are accessible to the public
	workDir, _ := os.Getwd()
	fileserver(r, "/images", http.Dir(filepath.Join(workDir, "/images")))
	fileserver(r, "/css", http.Dir(filepath.Join(workDir, "/css")))
	fileserver(r, "/js", http.Dir(filepath.Join(workDir, "/js")))

	// Start the server
	http.ListenAndServe(":3000", r)
}

// "fileserver conveniently sets up a http.FileServer handler to serve static files from a http.FileSystem."
// See https://github.com/go-chi/chi/blob/18d990c0d1c023b05a3652d322ae36d8bdb62e07/_examples/fileserver/main.go
func fileserver(r chi.Router, path string, root http.FileSystem) {
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
