package db

import (
	"net/http"

	auth "github.com/bokwoon95/orbital/auth"
	erro "github.com/bokwoon95/orbital/erro"
)

// GetNavbarData lorem ipsum
func GetNavbarData(r *http.Request) (
	loggedIn bool,
	uid int,
	displayName string,
	role string,
	participantTeamStatus string,
	err error,
) {
	var hashedCookie string
	hashedCookie, uid, err = auth.GetActiveSession(r)
	if hashedCookie != "" {
		loggedIn = true
	}

	var user User
	if loggedIn {
		user, err = GetUserByID(uid)
		erro.WrapX(&err)

		displayName = user.DisplayName
		role, err = GetRoleByID(uid)
		erro.WrapX(&err)
	}

	if role == Roles.Participant {
		participantTeamStatus, err = GetTeamStatusByID(uid)
		erro.WrapX(&err)
	}
	return
}
