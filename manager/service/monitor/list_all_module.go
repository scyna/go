package monitor

import (
	"github.com/scylladb/gocqlx/v2/qb"
	proto "github.com/scyna/go/manager/.proto/generated"
	"github.com/scyna/go/manager/model"
	"github.com/scyna/go/scyna"
)

func ListAllModule(s *scyna.Service, request *proto.ListAllModuleRequest) {

	var modules []*model.Module
	if err := qb.Select("scyna.module").
		Columns("code", "description", "org_code").
		Query(scyna.DB).
		SelectRelease(&modules); err != nil {
		s.Logger.Info(err.Error())
		s.Error(scyna.SERVER_ERROR)
		return
	}

	var response proto.ListAllModuleResponse
	for _, m := range modules {
		response.Items = append(response.Items, m.ToDTO())
	}

	s.Done(&response)
}
