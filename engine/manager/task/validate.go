package task

import (
	"errors"
	"regexp"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/scyna/go/scyna"
)

func validateAddTaskRequest(request *scyna.AddTaskRequest) error {
	if request.Time < time.Now().UnixNano() {
		return errors.New("Task time is less than now")
	}
	// TODO: check valid each field in request
	return nil
}

func validateAddRecurringTaskRequest(request *scyna.AddRecurringTaskRequest) error {
	if request.Time < time.Now().UnixNano() {
		return errors.New("Task time is less than now")
	}
	// TODO: check valid each field in request
	return nil
}

func validateCancelTaskRequest(request *scyna.CancelTaskRequest) error {
	return validation.ValidateStruct(request,
		validation.Field(&request.TaskID, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{27}$"))),
	)
}

func validateCancelRecurringTaskRequest(request *scyna.CancelRecurringTaskRequest) error {
	return validation.ValidateStruct(request,
		validation.Field(&request.TaskID, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{10,19}$"))),
	)
}
