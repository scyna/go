package monitor

import (
	"github.com/scylladb/gocqlx/v2/qb"
	proto "github.com/scyna/go/manager/.proto/generated"
	"github.com/scyna/go/scyna"
	"log"
	"time"
)

func ProcessSessionActive(s *scyna.Service, request *proto.ProcessSessionActiveRequest) {
	s.Logger.Info("Receive ProcessSessionActive request")

	var modules []module
	if err := qb.Select("scyna.module").
		Columns("code", "description").
		Query(scyna.DB).
		SelectRelease(&modules); err != nil {
		s.Logger.Info(err.Error())
		s.Error(scyna.SERVER_ERROR)
		return
	}

	log.Println(len(modules))

	var response proto.ProcessSessionActiveResponse
	point := time.Now().Add(-time.Minute * 5)
	for _, m := range modules {
		var sessions []session
		if err := qb.Select("scyna.session").
			Columns("id", "start", "last_update").
			Where(qb.GtOrEq("last_update")).
			AllowFiltering().
			Query(scyna.DB).
			Bind(point).
			SelectRelease(&sessions); err != nil {
			s.Logger.Info(err.Error())
			s.Error(scyna.SERVER_ERROR)
			return
		}
		item := &proto.ProcessSessionActiveResponse_Module{
			Code:     m.Code,
			Actives:  int32(len(sessions)),
			Sessions: nil,
		}

		for _, s := range sessions {
			item.Sessions = append(item.Sessions, &proto.ProcessSessionActiveResponse_Session{
				Id:         s.ID,
				Start:      s.Start.Format(time.RFC3339),
				LastUpdate: s.LastUpdate.Format(time.RFC3339),
			})
		}

		response.Items = append(response.Items, item)
	}

	s.Done(&response)
}

type module struct {
	Code        string `db:"code"`
	Description string `db:"description"`
}

type session struct {
	ID         string    `db:"id"`
	Start      time.Time `db:"start"`
	LastUpdate time.Time `db:"last_update"`
}
