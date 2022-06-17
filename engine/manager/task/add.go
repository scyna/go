package task

import (
	"fmt"
	"time"

	"github.com/gocql/gocql"
	"github.com/scyna/go/scyna"
)

const MAX_TRY_ADD_TASK = 3

func AddTask(s *scyna.Service, request *scyna.AddTaskRequest) {
	// Check validate task request (time > now, ...)
	if err := validateAddTaskRequest(request); err != nil {
		s.Error(scyna.REQUEST_INVALID)
		s.Logger.Error(err.Error())
		return
	}

	// Insert new task to scyna.task table
	var task = Task{
		ID:        scyna.ID.Next(),
		Topic:     request.Topic,
		Start:     time.Unix(request.Start, 0),
		Interval:  request.Interval,
		LoopCount: 0,
		Next:      time.Unix(request.Start, 0),
		LoopMax:   request.LoopMax,
		Data:      request.Data,
		Active:    true,
	}

	qBatch := scyna.DB.NewBatch(gocql.LoggedBatch)
	qBatch.Query("INSERT INTO scyna.task(id, topic, data, start, next, interval, loop_count, loop_max, active) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);",
		task.ID, task.Topic, task.Data, task.Start, task.Next, task.Interval, task.LoopCount, task.LoopMax, task.Active)

	qBatch.Query("INSERT INTO scyna.module_has_task(module, task_id) VALUES (?, ?);", request.Module, task.ID)

	// Generate period id
	// A group task contain all task must execute in a block 1 minute
	bucket := scyna.GetMinuteByTime(task.Start)
	qBatch.Query("INSERT INTO scyna.todo(bucket, task_id) VALUES (?, ?);", bucket, task.ID)
	if err := scyna.DB.ExecuteBatch(qBatch); err != nil {
		s.Error(scyna.REQUEST_INVALID)
		s.Logger.Error(err.Error())
		return
	}

	var response = scyna.AddTaskResponse{
		TaskID: fmt.Sprintf("%d", task.ID),
	}
	s.Done(&response)
}
