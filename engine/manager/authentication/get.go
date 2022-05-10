package authentication

import (
	"log"
	"time"

	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scyna/go/scyna"
)

func Get(s *scyna.Service, request *scyna.GetAuthRequest) {
	log.Println("Receive GetAuthRequest")
	if expired, userID := getAuthentication(request.Token, request.App); expired != nil {
		s.Done(&scyna.GetAuthResponse{
			Token:   request.Token,
			UserID:  userID,
			Expired: uint64(expired.UnixMicro()),
		})
	} else {
		s.LOG.Warning("Not exists Token, App")
		s.Error(scyna.REQUEST_INVALID)
		return
	}
}

func getAuthentication(token string, app string) (*time.Time, string) {
	/*check authentication*/
	var auth struct {
		Expired time.Time `db:"expired"`
		Apps    []string  `db:"apps"`
		UserID  string    `db:"uid"`
	}

	if err := qb.Select("scyna.authentication").
		Columns("expired", "apps", "uid").
		Where(qb.Eq("id")).
		Limit(1).
		Query(scyna.DB).
		Bind(token).
		GetRelease(&auth); err != nil {
		log.Println("authentication", err.Error())
		return nil, ""
	}

	hasApp := false
	for _, a := range auth.Apps {
		if a == app {
			hasApp = true
			break
		}
	}

	if !hasApp {
		log.Print("No app")
		return nil, ""
	}

	return &auth.Expired, auth.UserID
}
