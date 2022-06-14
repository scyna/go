package task

import "github.com/scyna/go/scyna"

func AddTask(s *scyna.Service, request *scyna.AddTaskRequest) {
	// TODO: check validate task request (time > now, ...)
	// TODO: insert new task to scyna.task table
}
