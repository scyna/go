package task

import (
	"errors"
	"regexp"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/scyna/go/scyna"
)

func validateAddTaskRequest(request *scyna.AddTaskRequest) error {
	if request.Start < time.Now().Unix() {
		return errors.New("Task time is less than now")
	}

	return validation.ValidateStruct(request,
		validation.Field(&request.Topic, validation.Required, validation.Length(1, 100)),
		validation.Field(&request.Interval, validation.Required, validation.Min(0)),
		validation.Field(&request.Module, validation.Required, validation.Length(1, 30)),
	)
}

func validateCancelTaskRequest(request *scyna.CancelTaskRequest) error {
	return validation.ValidateStruct(request,
		validation.Field(&request.TaskID, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{10,19}$"))),
	)
}
