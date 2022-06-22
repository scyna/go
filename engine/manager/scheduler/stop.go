package scheduler

import (
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scyna/go/scyna"
)

func StopTask(s *scyna.Service, request *scyna.StopTaskRequest) {

	if err := qb.Update("scyna.task").
		Set("done").
		Where(qb.Eq("id")).
		Query(scyna.DB).
		Bind(true, request.Id).
		ExecRelease(); err != nil {
		s.Error(scyna.REQUEST_INVALID)
		s.Logger.Error(err.Error())
		return
	}

	s.Done(scyna.OK)
}
