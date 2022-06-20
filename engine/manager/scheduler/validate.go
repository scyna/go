package scheduler

import (
	"errors"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/scyna/go/scyna"
)

func validateStartTaskRequest(request *scyna.StartTaskRequest) error {
	if int64(request.Time) < time.Now().Unix() {
		return errors.New("Task time is less than now")
	}
	if request.Interval < 60 {
		return errors.New("interval must be greater than 60 second")
	}
	return validation.ValidateStruct(request,
		validation.Field(&request.Topic, validation.Required, validation.Length(1, 100)),
		validation.Field(&request.Module, validation.Required, validation.Length(1, 30)),
	)
}

func validateStopTaskRequest(request *scyna.StopTaskRequest) error {
	if request.Id == 0 {
		return errors.New("invalid task id")
	}
	return nil
}
