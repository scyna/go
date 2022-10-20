package setting

import (
	"log"

	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scyna/go/scyna"
)

func Read(s *scyna.Service, request *scyna.ReadSettingRequest) {
	log.Println("Receive ReadSettingRequest")

	var value string
	if err := qb.Select("scyna.setting").
		Columns("value").
		Where(qb.Eq("context"), qb.Eq("key")).
		Limit(1).
		Query(scyna.DB).
		Bind(request.Context, request.Key).
		GetRelease(&value); err != nil {
		s.Logger.Info("Can not read setting - " + err.Error())
		s.Error(scyna.REQUEST_INVALID)
		return
	}

	s.Done(&scyna.ReadSettingResponse{Value: value})
}
