package task

import (
	"errors"
	"time"

	"github.com/scyna/go/scyna"
)

func validateAddTaskRequest(request *scyna.AddTaskRequest) error {
	if request.Time < time.Now().UnixNano() {
		return errors.New("Task time is less than now")
	}
	return nil
}

func validateAddRecurringTaskRequest(request *scyna.AddRecurringTaskRequest) error {
	if request.Time < time.Now().UnixNano() {
		return errors.New("Task time is less than now")
	}
	return nil
}
