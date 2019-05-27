package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	auth "github.com/bokwoon95/orbital/auth"
	controllers "github.com/bokwoon95/orbital/orbital_controllers"
	db "github.com/bokwoon95/orbital/orbital_db"
	"github.com/go-chi/chi"
)

func main() {
	// configure log to print in format '<date> <time> <filename:linenumber> message'
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Initialize database and defer its closing to the end of main()
	err := db.Init()
	if err != nil {
		log.Fatalln(err)
	}
	defer db.DB.Close()

	// Initialize session database
	err = auth.Init()
	if err != nil {
		log.Fatalln(err)
	}
	defer auth.SessionDB.Close()

	// Initialize router
	r := chi.NewRouter()

	// Set up routes
	r.Get("/testdb", controllers.TestDB)
	r.Get("/", controllers.AboutOrbital)
	r.Get("/showcase", controllers.PastYearShowcase)
	r.Get("/register", controllers.RegisterGet)
	r.Post("/register", controllers.RegisterPost)
	r.Get("/login", controllers.LoginGet)
	r.Post("/login", controllers.LoginPost)
	r.Post("/logout", controllers.Logout)

	r.Get("/cookie", controllers.CookieInspectorGet)
	r.Post("/cookie", controllers.CookieInspectorPost)
	r.Get("/session", controllers.SessionGet)
	r.Post("/sessionset", controllers.SessionSet)
	r.Post("/sessionrevoke", controllers.SessionRevoke)

	// Ensure the below directories are accessible to the public
	workDir, _ := os.Getwd()
	fileserver(r, "/orbital_views", http.Dir(filepath.Join(workDir, "/orbital_views")))

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
