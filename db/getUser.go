package db

import "database/sql"

// User lorem ipsum
type User struct {
	ID           int            `db:"uid"`
	NUSNETID     string         `db:"nusnetid"`
	PasswordHash string         `db:"password"`
	DisplayName  string         `db:"display_name"`
	OpenID       sql.NullString `db:"openid"`
	Email        sql.NullString `db:"email"`
}

// GetUserByID lorem ipsum
func GetUserByID(uid int) (User, error) {
	var user User
	err := DB.QueryRowx(`
	SELECT uid, nusnetid, display_name, password, display_name, openid, email
	FROM users WHERE uid = $1
	`, uid).StructScan(&user)
	return user, err
}

// GetUserByNUSNET lorem ipsum
func GetUserByNUSNET(nusnetid string) (User, error) {
	var user User
	err := DB.QueryRowx(`
	SELECT uid, nusnetid, display_name, password, openid, email
	FROM users WHERE nusnetid = $1
	`, nusnetid).StructScan(&user)
	return user, err
}
