package db

import (
	erro "github.com/bokwoon95/orbital/erro"
)

// ParticipantTeamStatusesStruct lorem ipsum
type ParticipantTeamStatusesStruct struct {
	Teamed        string
	Teamless      string
	InviteSent    string
	InvitePending string
}

// ParticipantTeamStatuses lorem ipsum
// DO NOT modify this struct variable anywhere! It is intended to be a
// constant, but golang does not have struct constants so it's up to you, dear
// reader, to not modify it
var ParticipantTeamStatuses = ParticipantTeamStatusesStruct{
	Teamed:        "teamed",
	Teamless:      "teamless",
	InviteSent:    "invitesent",
	InvitePending: "invitepending",
}

// GetTeamStatusByID lorem ipsum
func GetTeamStatusByID(uid int) (string, error) {
	var participantTeamStatus string
	var teamed bool
	err := DB.QueryRowx("SELECT EXISTS(SELECT 1 FROM participants WHERE uid = $1 AND team IS NOT NULL);", uid).Scan(&teamed)
	erro.WrapX(&err)
	if teamed {
		participantTeamStatus = ParticipantTeamStatuses.Teamed
	} else {
		participantTeamStatus = ParticipantTeamStatuses.Teamless
	}
	return participantTeamStatus, err
}
