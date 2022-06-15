package task

import (
	"fmt"

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
		Time:     request.Time,
		Interval: request.Interval,
		SendTo:   request.To,
		Type:     request.Type,
		Data:     request.Data,
	}
	// Insert task to scyna.recurring_task table
	for i := 0; i < MAX_TRY_ADD_TASK; i++ {
		if applied, err := qb.Insert("scyna.recurring_task").
			Columns("period", "time", "interval", "send_to", "type", "data").
			Unique().
			Query(scyna.DB).
			BindStruct(task).
			ExecCASRelease(); !applied {
			if err != nil {
				s.Error(scyna.SERVER_ERROR)
				scyna.LOG.Error(err.Error())
				return
			} else {
				task.Time += random.Int63n(1000000000)
			}
		}
	}

	// TODO: Insert task to scyna.task table
	var response = scyna.AddRecurringTaskResponse{
		TaskID: fmt.Sprintf("%015d", request.Time),
	}

	s.Done(&response)
}
