package auth

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"hash"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // postgres driver for sqlx
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

var (
	// SessionDB lorem ipsum
	SessionDB   *sqlx.DB
	hmacEncoder hash.Hash
)

const hmacSecretKey = "hmac-secret-key" // should be set as env var in the future
const cookieName = "rawCookie"

// Init will initialize the SessionDB object for storing sessions
func Init() {
	var err error
	if SessionDB, err = sqlx.Open(
		"postgres",
		"postgres://bokwoon@localhost:5434/sessiondb_dev?sslmode=disable",
	); err != nil {
		panic(err)
	}
	hmacEncoder = hmac.New(sha256.New, []byte(hmacSecretKey))
}

// HashPassword returns a hash of the user's password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CompareHashAndPassword compares a hash from the db and the password provided by the user
func CompareHashAndPassword(passwordhash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordhash), []byte(password))
}

func generateRandomString() (string, error) {
	arr := make([]byte, 32)
	_, err := rand.Read(arr)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(arr), nil
}

func hashCookie(rawCookie string) string {
	hmacEncoder.Reset()
	hmacEncoder.Write([]byte(rawCookie))
	b := hmacEncoder.Sum(nil)
	return base64.URLEncoding.EncodeToString(b)
}

// SetSession lorem ipsum
func SetSession(w http.ResponseWriter, r *http.Request, userid int) error {
	rawCookie, err := generateRandomString()
	if err != nil {
		return errors.Wrap(err, "failed creating random string for rawCookie")
	}
	hashedCookie := hashCookie(rawCookie)
	http.SetCookie(w, &http.Cookie{
		Name:     cookieName,
		Value:    rawCookie,
		HttpOnly: true, // disable JavaScript access to cookie
		// Secure:   true, // allow sending only over HTTPS
		// SameSite: http.SameSiteStrictMode,
		MaxAge: int((time.Hour * 24 * 30).Seconds()), // one month (in seconds)
		// Path: "",
	})
	_, err = SessionDB.Exec("INSERT INTO sessions (hash, id) VALUES ($1, $2)", hashedCookie, userid)
	if err != nil {
		return errors.Wrap(err, "failed to insert hash into sessions table")
	}
	return nil
}

// GetActiveSession lorem ipsum
func GetActiveSession(r *http.Request) (string, int) {
	rawCookie, err := r.Cookie(cookieName)
	if err != nil {
		return "", 0
	}
	hashedCookie := hashCookie(rawCookie.Value)
	fmt.Println("hashedCookie is " + hashedCookie)
	session := &struct {
		Hash string `db:"hash"`
		Id   int    `db:"id"`
	}{}
	err = SessionDB.QueryRowx("SELECT hash, id FROM sessions WHERE hash = $1 LIMIT 1", hashedCookie).StructScan(session)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("SELECT hash, id FROM sessions WHERE hash = " + hashedCookie)
	fmt.Println("session.hash:" + session.Hash + " session.id:" + string(session.Id))
	return session.Hash, session.Id
}

// RevokeSession will wipe the session cookie of anyone who ends up calling it
func RevokeSession(w http.ResponseWriter, r *http.Request) error {
	rawCookie, err := r.Cookie(cookieName)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("cookie '%s' does not exist in user's browser", cookieName))
	}
	hashedCookie := hashCookie(rawCookie.Value)
	http.SetCookie(w, &http.Cookie{
		Name:   cookieName,
		Value:  "",
		MaxAge: 0,
		Path:   "/",
	})
	_, err = SessionDB.Exec("DELETE FROM sessions WHERE hash = $1", hashedCookie)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed trying to delete session row from sessions table"))
	}
	return nil
}

// RefreshSession will reissue
// func RefreshSession(w http.ResponseWriter, r *http.Request) {
// 	RevokeSession(w, r)
// }
