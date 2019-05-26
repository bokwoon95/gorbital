package db

import (
	"github.com/bokwoon95/orbital/erro"
)

// RolesStruct lorem ipsum
type RolesStruct struct {
	Empty       string
	Participant string
	Tutor       string
	Adviser     string
	Mentor      string
	Admin       string
}

// Roles lorem ipsum
// DO NOT modify this struct variable anywhere! It is intended to be a
// constant, but golang does not have struct constants so it's up to you, dear
// reader, to not modify it
var Roles = RolesStruct{
	Empty:       "",
	Participant: "participant",
	Tutor:       "tutor",
	Adviser:     "adviser",
	Mentor:      "mentor",
	Admin:       "admin",
}

// GetRoleByID lorem ipsum
func GetRoleByID(uid int) (string, error) {
	var exists bool
	var err error
	for _, role := range []string{
		Roles.Participant,
		Roles.Tutor,
		Roles.Adviser,
		Roles.Mentor,
		Roles.Admin,
	} {
		tablename := string(role) + "s"
		err = DB.QueryRowx("SELECT EXISTS(SELECT 1 FROM "+tablename+" WHERE uid = $1)", uid).Scan(&exists)
		if err != nil {
			return Roles.Empty, erro.Wrap(err)
		} else if exists {
			return role, nil
		}
	}
	return Roles.Empty, nil
}
