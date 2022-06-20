package scheduler

import (
	"github.com/scyna/go/scyna"
)

func StopTask(s *scyna.Service, request *scyna.StopTaskRequest) {
	// Check validate params
	if err := validateStopTaskRequest(request); err != nil {
		s.Error(scyna.REQUEST_INVALID)
		s.Logger.Error(err.Error())
		return
	}
	// Mark active = false
	var task = Task{
		ID: request.ID,
	}
	if err := task.Deactive(); err != nil {
		s.Error(scyna.REQUEST_INVALID)
		s.Logger.Error(err.Error())
		return
	}

	// Done
	s.Done(scyna.OK)
}
