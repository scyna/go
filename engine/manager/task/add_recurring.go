package task

import (
	"fmt"
	"time"

	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scyna/go/scyna"
)

func AddRecurringTask(s *scyna.Service, request *scyna.AddRecurringTaskRequest) {
	// Check validate task
	if err := validateAddRecurringTaskRequest(request); err != nil {
		s.Error(scyna.REQUEST_INVALID)
		return
	}

	var task = RecurringTask{
		Time:     time.Unix(0, request.Time),
		Interval: request.Interval,
		SendTo:   request.SendTo,
		Type:     request.Type,
		Data:     request.Data,
		Count:    request.Count,
		ID:       scyna.ID.Next(),
	}

	// Insert task to scyna.recurring_task table
	if applied, err := qb.Insert("scyna.recurring_task").
		Columns("id", "time", "interval", "count", "send_to", "type", "data").
		Unique().
		Query(scyna.DB).
		BindStruct(task).
		ExecCASRelease(); !applied {
		if err != nil {
			scyna.LOG.Error(err.Error())
		} else {
			scyna.LOG.Error(fmt.Sprintf("Insert recurring task not applied: %+v\n", task))
		}
		s.Error(scyna.SERVER_ERROR)
		return
	}

	// Insert task to scyna.task table
	var addTaskResponse scyna.AddTaskResponse
	if err := s.CallService(scyna.ADD_TASK_URL, &scyna.AddTaskRequest{
		Time:            request.Time,
		RecurringTaskID: task.ID,
		Type:            task.Type,
		SendTo:          task.SendTo,
		Data:            task.Data,
	}, &addTaskResponse); err.Code != scyna.OK.Code {
		s.Error(scyna.SERVER_ERROR)
		// TODO: roll back in scyna.recurring_task table
		return
	}

	var response = scyna.AddRecurringTaskResponse{
		TaskID: fmt.Sprintf("%d", task.ID),
	}

	s.Done(&response)
}
