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

	"github.com/bokwoon95/orbital/erro"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // postgres driver for sqlx
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
func Init() error {
	var err error
	if SessionDB, err = sqlx.Open(
		"postgres",
		"postgres://bokwoon@localhost:5434/sessiondb_dev?sslmode=disable",
	); err != nil {
		return erro.Wrap(err)
	}
	hmacEncoder = hmac.New(sha256.New, []byte(hmacSecretKey))
	return nil
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
		return "", erro.Wrap(err)
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
		return erro.Wrap(err)
	}
	hashedCookie := hashCookie(rawCookie)
	_, err = SessionDB.Exec("INSERT INTO sessions (hash, id) VALUES ($1, $2)", hashedCookie, userid)
	if err != nil {
		return erro.Wrap(err)
	}
	http.SetCookie(w, &http.Cookie{
		Name:     cookieName,
		Value:    rawCookie,
		HttpOnly: true, // disable JavaScript access to cookie
		// Secure:   true, // allow sending only over HTTPS
		// SameSite: http.SameSiteStrictMode,
		MaxAge: int((time.Hour * 24 * 30).Seconds()), // one month (in seconds)
		// Path: "",
	})
	return nil
}

// GetActiveSession returns the hashedCookie and uid of the current logged in
// user. If nobody is logged in or if the user's rawCookie is not present in
// the 'sessions' table, it returns an empty string and 0
func GetActiveSession(r *http.Request) (string, int, error) {
	rawCookieObj, err := r.Cookie(cookieName)
	if err != nil {
		if err == http.ErrNoCookie {
			return "", 0, nil
		}
		return "", 0, erro.Wrap(err)
	}
	hashedCookie := hashCookie(rawCookieObj.Value)
	session := struct {
		Hash string `db:"hash"`
		ID   int    `db:"id"`
	}{}
	err = SessionDB.QueryRowx("SELECT hash, id FROM sessions WHERE hash = $1 LIMIT 1", hashedCookie).StructScan(&session)
	if err != nil {
		fmt.Println(err)
	}
	return session.Hash, session.ID, nil
}

// RevokeSession will wipe the session cookie of anyone who ends up calling it
func RevokeSession(w http.ResponseWriter, r *http.Request) error {
	hashedCookie, _, _ := GetActiveSession(r)
	http.SetCookie(w, &http.Cookie{
		Name:   cookieName,
		Value:  "",
		MaxAge: 0,
		Path:   "/",
	})
	_, err := SessionDB.Exec("DELETE FROM sessions WHERE hash = $1", hashedCookie)
	if err != nil {
		return erro.Wrap(err)
	}
	return nil
}
