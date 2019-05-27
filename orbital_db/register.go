package db

import erro "github.com/bokwoon95/orbital/erro"

// InsertParticipant lorem ipsum
func InsertParticipant(nusnetid string, passwordhash string, displayname string) (int, error) {
	var uid int
	err := DB.QueryRowx(`
INSERT INTO users (nusnetid, password, display_name)
VALUES ($1, $2, $3) RETURNING uid
	`, nusnetid, passwordhash, displayname).Scan(&uid)
	if err != nil {
		return 0, erro.Wrap(err)
	}
	_, err = DB.Exec(`INSERT INTO participants (uid) VALUES ($1)`, uid)
	if err != nil {
		return 0, erro.Wrap(err)
	}
	return uid, nil
}
