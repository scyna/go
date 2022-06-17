package task

import (
	"strconv"

	"github.com/gocql/gocql"
	"github.com/scyna/go/scyna"
)

func CancelRecurringTask(s *scyna.Service, request *scyna.CancelRecurringTaskRequest) {
	// check validation of request
	if err := validateCancelRecurringTaskRequest(request); err != nil {
		s.Error(scyna.REQUEST_INVALID)
		s.Logger.Error(err.Error())
		return
	}

	taskID, _ := strconv.ParseUint(request.TaskID, 10, 64)
	// remove in scyna.task and scyna.recurring_task
	qBatch := scyna.DB.NewBatch(gocql.LoggedBatch)
	qBatch.Query("DELETE FROM scyna.recurring_task WHERE id = ?;", taskID)
	qBatch.Query("DELETE FROM scyna.task WHERE recurring_task_id = ? ALLOW FILTERING", taskID)

	// Remove in scyna.task
	if err := scyna.DB.ExecuteBatch(qBatch); err != nil {
		s.Error(scyna.REQUEST_INVALID)
		s.Logger.Error(err.Error())
		return
	}
	s.Done(scyna.OK)
}
