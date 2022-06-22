package scheduler

import (
	"errors"
	"math"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gocql/gocql"
	"github.com/scyna/go/scyna"
)

func StartTask(s *scyna.Service, request *scyna.StartTaskRequest) {
	if err := validateStartTaskRequest(request); err != nil {
		s.Error(scyna.REQUEST_INVALID)
		s.Logger.Error(err.Error())
		return
	}

	if request.Loop == 0 {
		request.Loop = math.MaxInt64
	}

	// Insert new task to scyna.task table
	taskID := scyna.ID.Next()
	start := time.Unix(int64(request.Time), 0)
	qBatch := scyna.DB.NewBatch(gocql.LoggedBatch)
	qBatch.Query("INSERT INTO scyna.task(id, topic, data, start, next, interval, loop_count, loop_index, done) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);",
		taskID, request.Topic, request.Data, start, start, request.Interval, request.Loop, 0, false)

	qBatch.Query("INSERT INTO scyna.module_has_task(module, task_id) VALUES (?, ?);", request.Module, taskID)

	bucket := GetBucket(start) // Generate period id
	qBatch.Query("INSERT INTO scyna.todo(bucket, task_id) VALUES (?, ?);", bucket, taskID)
	if err := scyna.DB.ExecuteBatch(qBatch); err != nil {
		s.Error(scyna.REQUEST_INVALID)
		s.Logger.Error(err.Error())
		return
	}

	s.Done(&scyna.StartTaskResponse{Id: taskID})
}

func validateStartTaskRequest(request *scyna.StartTaskRequest) error {
	if int64(request.Time) < time.Now().Unix() {
		return errors.New("Task time is less than now")
	}
	if request.Interval < 60 {
		return errors.New("interval must be greater than 60 second")
	}
	return validation.ValidateStruct(request,
		validation.Field(&request.Topic, validation.Required, validation.Length(1, 255)),
		validation.Field(&request.Module, validation.Required, validation.Length(1, 255)),
	)
}
