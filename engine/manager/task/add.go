package task

import (
	"math/rand"
	"time"

	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scyna/go/scyna"
)

const MAX_TRY_ADD_TASK = 3

func AddTask(s *scyna.Service, request *scyna.AddTaskRequest) {
	// check validate task request (time > now, ...)
	if err := validateAddTaskRequest(request); err != nil {
		s.Error(scyna.REQUEST_INVALID)
		return
	}
	// Generate period id
	// A group task contain all task must execute in a block 2 minute
	period := scyna.GetMinuteByTime(time.Unix(0, request.Time)) / 2
	// Insert new task to scyna.task table
	var task = Task{
		Period: period,
		Time:   request.Time,
		Repeat: request.Repeat,
		SendTo: request.To,
		Type:   request.Type,
		Data:   request.Data,
	}

	for i := 0; i < MAX_TRY_ADD_TASK; i++ {
		if applied, err := qb.Insert("scyna.task").
			Columns("period", "time", "repeat", "context").
			Unique().
			Query(scyna.DB).
			BindStruct(task).
			ExecCASRelease(); !applied {
			if err != nil {
				s.Error(scyna.SERVER_ERROR)
				scyna.LOG.Error(err.Error())
				return
			} else {
				task.Time += rand.Int63n(100000000)
				continue
			}
		} else {
			break
		}
		s.Error(scyna.SERVER_ERROR)
		scyna.LOG.Error("Exceeding max times insert task")
		return
	}

	var response scyna.AddTaskResponse
	s.Done(&response)
}
