package scheduler

import (
	"time"

	"github.com/gocql/gocql"
	"github.com/scyna/go/scyna"
)

func StartTask(s *scyna.Service, request *scyna.StartTaskRequest) {
	// Check validate task request (time > now, ...)
	if err := validateStartTaskRequest(request); err != nil {
		s.Error(scyna.REQUEST_INVALID)
		s.Logger.Error(err.Error())
		return
	}

	// Insert new task to scyna.task table
	taskID := scyna.ID.Next()
	var task = Task{
		ID:        taskID,
		Topic:     request.Topic,
		Start:     time.Unix(int64(request.Time), 0),
		Interval:  request.Interval,
		LoopCount: request.Loop,
		Next:      time.Unix(int64(request.Time), 0),
		LoopIndex: 0,
		Data:      request.Data,
		Done:      false,
	}

	qBatch := scyna.DB.NewBatch(gocql.LoggedBatch)
	qBatch.Query("INSERT INTO scyna.task(id, topic, data, start, next, interval, loop_count, loop_index, done) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);",
		task.ID, task.Topic, task.Data, task.Start, task.Next, task.Interval, task.LoopCount, task.LoopIndex, task.Done)

	qBatch.Query("INSERT INTO scyna.module_has_task(module, task_id) VALUES (?, ?);", request.Module, task.ID)

	// Generate period id
	// A group task contain all task must execute in a block 1 minute
	bucket := GetBucket(task.Start)
	qBatch.Query("INSERT INTO scyna.todo(bucket, task_id) VALUES (?, ?);", bucket, task.ID)
	if err := scyna.DB.ExecuteBatch(qBatch); err != nil {
		s.Error(scyna.REQUEST_INVALID)
		s.Logger.Error(err.Error())
		return
	}

	var response = scyna.StartTaskResponse{
		Id: task.ID,
	}
	s.Done(&response)
}
