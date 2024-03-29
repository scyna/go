package setting

import (
	"fmt"
	"log"

	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scyna/go/scyna"
)

func Read(s *scyna.Service, request *scyna.ReadSettingRequest) {
	log.Println("Receive ReadSettingRequest")

	var value string
	if err := qb.Select("scyna.setting").
		Columns("value").
		Where(qb.Eq("module_code"), qb.Eq("key")).
		Limit(1).
		Query(scyna.DB).
		Bind(request.Module, request.Key).
		GetRelease(&value); err != nil {
		s.Logger.Info(fmt.Sprintf("Cannot read setting: %s - request = %s\n", err.Error(), request.String()))
		s.Error(scyna.REQUEST_INVALID)
		return
	}

	s.Done(&scyna.ReadSettingResponse{Value: value})
}
