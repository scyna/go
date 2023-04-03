package monitor

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/scylladb/gocqlx/v2/qb"
	proto "github.com/scyna/go/manager/.proto/generated"
	"github.com/scyna/go/manager/model"
	"github.com/scyna/go/scyna"
)

func GetActiveSession(s *scyna.Service, request *proto.ListActiveSessionRequest) {
	if validateGetActiveSession(request) != nil {
		s.Error(scyna.REQUEST_INVALID)
		return
	}

	var sessions []*model.Session

	if err := qb.Select("scyna.session").
		Columns("id", "module_code", "start", "last_update", "end", "exit_code").
		Where(qb.Eq("module_code"), qb.Gt("last_update")).
		AllowFiltering().
		Query(scyna.DB).
		Bind(request.ModuleCode, time.Now().Add(-time.Minute*6)).
		SelectRelease(&sessions); err != nil {
		s.Logger.Info(err.Error())
		s.Error(scyna.SERVER_ERROR)
		return
	}

	var response proto.ListSessionResponse
	for _, s := range sessions {
		response.Items = append(response.Items, s.ToDTO())
	}
	response.Total = uint32(len(sessions))
	s.Done(&response)
}

func validateGetActiveSession(request *proto.ListActiveSessionRequest) error {
	return validation.ValidateStruct(request,
		validation.Field(&request.ModuleCode, validation.Required, validation.Length(5, 100)))
}
