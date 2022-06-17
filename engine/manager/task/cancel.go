package task

import (
	"strconv"

	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scyna/go/scyna"
)

func CancelTask(s *scyna.Service, request *scyna.CancelTaskRequest) {
	// Check validate params
	if err := validateCancelTaskRequest(request); err != nil {
		s.Error(scyna.REQUEST_INVALID)
		s.Logger.Error(err.Error())
		return
	}
	// Parse task
	taskID, _ := strconv.ParseUint(request.TaskID, 10, 64)

	// Mark active = false
	if err := qb.Update("scyna.task").
		Set("active").
		Where(qb.Eq("id")).
		Query(scyna.DB).
		Bind(false, taskID).
		ExecRelease(); err != nil {
		s.Error(scyna.REQUEST_INVALID)
		s.Logger.Error(err.Error())
		return
	}

	// Done
	s.Done(scyna.OK)
}
