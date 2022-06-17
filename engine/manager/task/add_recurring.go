package task

import (
	"fmt"
	"time"

	"github.com/gocql/gocql"
	"github.com/scyna/go/scyna"
)

func AddRecurringTask(s *scyna.Service, request *scyna.AddRecurringTaskRequest) {
	// Check validate task
	if err := validateAddRecurringTaskRequest(request); err != nil {
		s.Error(scyna.REQUEST_INVALID)
		return
	}

	var rTask = RecurringTask{
		Time:     time.Unix(0, request.Time),
		Interval: request.Interval,
		SendTo:   request.SendTo,
		Type:     request.Type,
		Data:     request.Data,
		Count:    request.Count,
		ID:       scyna.ID.Next(),
	}
	// Insert task to scyna.recurring_task table
	qBatch := scyna.DB.NewBatch(gocql.LoggedBatch)
	qBatch.Query("INSERT INTO scyna.recurring_task (id, time, interval, send_to, type, data, count) VALUES (?, ?, ?, ?, ?, ?, ?);",
		rTask.ID, rTask.Time, rTask.Interval, rTask.SendTo, rTask.Type, rTask.Data, rTask.Count)
	// A group task contain all task must execute in a block 1 minute
	bucket := scyna.GetMinuteByTime(time.Unix(0, request.Time))
	var task = Task{
		Bucket:          bucket,
		ID:              scyna.ID.Next(),
		Time:            rTask.Time,
		RecurringTaskID: rTask.ID,
		SendTo:          rTask.SendTo,
		Type:            rTask.Type,
		Data:            rTask.Data,
	}
	// Insert to scyna.task table
	qBatch.Query("INSERT INTO scyna.task (bucket, id, time, recurring_task_id, send_to, type, data) VALUES (?, ?, ?, ?, ?, ?, ?);",
		task.Bucket, task.ID, task.Time, task.RecurringTaskID, task.SendTo, task.Type, task.Data)

	if err := scyna.DB.ExecuteBatch(qBatch); err != nil {
		s.Error(scyna.REQUEST_INVALID)
		s.Logger.Error(err.Error())
		return
	}

	// Response
	var response = scyna.AddRecurringTaskResponse{
		TaskID: fmt.Sprintf("%d", rTask.ID),
	}
	s.Done(&response)
}
