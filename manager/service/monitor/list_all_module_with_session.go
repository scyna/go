package monitor

import (
	"strconv"
	"time"

	"github.com/scylladb/gocqlx/v2/qb"
	proto "github.com/scyna/go/manager/.proto/generated"
	"github.com/scyna/go/manager/model"
	"github.com/scyna/go/scyna"
)

func ListAllModuleWithSession(s *scyna.Service, request *proto.ListAllModuleRequest) {

	var modules []*model.Module
	if err := qb.Select("scyna.module").
		Columns("code", "description", "").
		Query(scyna.DB).
		SelectRelease(&modules); err != nil {
		s.Logger.Info(err.Error())
		s.Error(scyna.SERVER_ERROR)
		return
	}

	s.Logger.Info("Modules: " + strconv.Itoa((len(modules))))

	var response proto.ListAllModuleResponse
	point := time.Now().Add(-time.Minute * 6)
	for _, m := range modules {
		var sessions []*model.Session
		if err := qb.Select("scyna.session").
			Columns("id", "module_code", "start", "last_update", "end", "exit_code").
			Where(qb.Eq("module_code"), qb.Gt("last_update")).
			AllowFiltering().
			Query(scyna.DB).
			Bind(m.Code, point).
			SelectRelease(&sessions); err != nil {
			s.Logger.Info(err.Error())
			s.Error(scyna.SERVER_ERROR)
			return
		}
		var item *proto.Module
		item.Code = m.Code
		item.TotalActive = uint32(len(sessions))

		for _, s := range sessions {
			item.Sessions = append(item.Sessions, s.ToDTO())
		}

		response.Items = append(response.Items, item)
		s.Logger.Info("Module: " + m.Code + " - " + strconv.Itoa((len(sessions))))
	}

	s.Done(&response)
}
