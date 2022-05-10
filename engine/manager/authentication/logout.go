package authentication

import (
	"log"
	"time"

	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scyna/go/scyna"
)

func Logout(s *scyna.Service, request *scyna.LogoutRequest) {
	log.Println("Receive LogoutRequest")

	if !checkOrg(request.Organization, request.Secret) {
		s.LOG.Warning("Organization not exist")
		s.Error(scyna.REQUEST_INVALID)
		return
	}

	if err := updateSession(request.Token, request.UserID); err != scyna.OK {
		s.Error(err)
		return
	}
	s.Done(scyna.OK)
}

func updateSession(token string, userID string) *scyna.Error {
	var userID_ string
	err := qb.Select("scyna.authentication").
		Columns("uid").
		Where(qb.Eq("id")).
		Limit(1).
		Query(scyna.DB).
		Bind(token).
		GetRelease(&userID_)
	if err != nil {
		log.Print("Error:", err)
		return scyna.SERVER_ERROR
	}

	if userID != userID_ {
		return scyna.SERVER_ERROR
	}

	now := time.Now()

	err = qb.Update("scyna.authentication").
		Set("expired").
		Where(qb.Eq("id")).
		Query(scyna.DB).
		Bind(now, token).
		ExecRelease()

	if err != nil {
		log.Print("Error:", err)
		return scyna.SERVER_ERROR
	}
	return nil
}
