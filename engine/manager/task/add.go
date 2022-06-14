package task

import "github.com/scyna/go/scyna"

const MAX_TRY_ADD_TASK = 3

func AddTask(s *scyna.Service, request *scyna.AddTaskRequest) {
	// TODO: check validate task request (time > now, ...)
	// TODO: insert new task to scyna.task table

	s.Done(scyna.OK)
}
