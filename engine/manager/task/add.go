package task

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scyna/go/scyna"
)

const MAX_TRY_ADD_TASK = 3

var random = rand.New(rand.NewSource(time.Now().UnixMicro()))

func AddTask(s *scyna.Service, request *scyna.AddTaskRequest) {
	// Check validate task request (time > now, ...)
	if err := validateAddTaskRequest(request); err != nil {
		s.Error(scyna.REQUEST_INVALID)
		return
	}
	// Generate period id
	// A group task contain all task must execute in a block 1 minute
	bucket := scyna.GetMinuteByTime(time.Unix(0, request.Time))
	// Insert new task to scyna.task table
	var task = Task{
		ID:              scyna.ID.Next(),
		Bucket:          bucket,
		Time:            time.Unix(0, request.Time),
		RecurringTaskID: request.RecurringTaskID,
		SendTo:          request.SendTo,
		Type:            request.Type,
		Data:            request.Data,
	}

	for i := 0; i < MAX_TRY_ADD_TASK; i++ {
		if applied, err := qb.Insert("scyna.task").
			Columns("bucket", "id", "recurring_task_id", "send_to", "type", "time", "data").
			Unique().
			Query(scyna.DB).
			BindStruct(task).
			ExecCASRelease(); !applied {
			if err != nil {
				s.Error(scyna.SERVER_ERROR)
				scyna.LOG.Error(err.Error())
				return
			} else {
				task.ID += scyna.ID.Next()
				continue
			}
		} else {
			break
		}
		s.Error(scyna.SERVER_ERROR)
		scyna.LOG.Error("Exceeding max times insert task")
		return
	}

	var response = scyna.AddTaskResponse{
		TaskID: fmt.Sprintf("%08d%019d", bucket, task.ID),
	}
	s.Done(&response)
}
