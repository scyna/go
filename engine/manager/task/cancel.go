package task

import (
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
	bucket := request.TaskID[:8]
	taskID := request.TaskID[8:]
	// Remove task in db
	if err := qb.Delete("scyna.task").
		Where(qb.Eq("bucket"), qb.Eq("id")).
		Query(scyna.DB).
		Bind(bucket, taskID).
		ExecRelease(); err != nil {
		s.Error(scyna.REQUEST_INVALID)
		s.Logger.Error(err.Error())
		return
	}

	// Done
	s.Done(scyna.OK)
}
