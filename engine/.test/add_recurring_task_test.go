package test

import (
	"testing"
	"time"

	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scyna/go/engine/manager/task"
	"github.com/scyna/go/scyna"
	scyna_test "github.com/scyna/go/scyna/testing"
)

func TestAddRecurringTask(t *testing.T) {
	var response scyna.AddRecurringTaskResponse
	scyna_test.ServiceTest(scyna.ADD_RECURRING_TASK_URL).
		WithRequest(&scyna.AddRecurringTaskRequest{
			Time:     time.Now().Add(time.Second * 10).UnixNano(),
			Interval: int64(time.Second * 60),
			Type:     "send-otp",
			SendTo:   "JETSTREAM.send_otp",
			Count:    10,
		}).ExpectSuccess().Run(t, &response)
	t.Logf("TaskID: %s\n", response.TaskID)

	/* DB */
	var rTask task.RecurringTask
	if err := qb.Select("scyna.recurring_task").
		Columns("*").
		Where(qb.Eq("id")).
		Query(scyna.DB).
		Bind(response.TaskID).
		GetRelease(&rTask); err != nil {
		t.Fatalf("Cannot get recurring task: %s\n", response.TaskID)
	}
	t.Logf("Recurring Task: %+v\n", rTask)
	defer qb.Delete("scyna.recurring_task").Where(qb.Eq("id")).Query(scyna.DB).Bind(response.TaskID).ExecRelease()

	var task task.Task
	if err := qb.Select("scyna.task").
		Columns("*").
		Where(qb.Eq("recurring_task_id")).
		AllowFiltering().
		Query(scyna.DB).
		Bind(response.TaskID).
		GetRelease(&task); err != nil {
		t.Fatalf("Cannot get task: %s\n", response.TaskID)
	}
	t.Logf("Task: %+v\n", task)
	defer qb.Delete("scyna.task").Where(qb.Eq("bucket"), qb.Eq("id")).Query(scyna.DB).Bind(task.Bucket, task.ID).ExecRelease()
}
