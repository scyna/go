package test

import (
	"testing"
	"time"

	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scyna/go/engine/manager/task"
	"github.com/scyna/go/scyna"
	scyna_test "github.com/scyna/go/scyna/testing"
)

func TestAddTask(t *testing.T) {
	var response scyna.AddTaskResponse
	scyna_test.ServiceTest(scyna.ADD_TASK_URL).WithRequest(&scyna.AddTaskRequest{
		Time:            time.Now().Add(time.Second * 10).UnixNano(),
		RecurringTaskID: 0,
		Type:            "send-otp",
		SendTo:          "JETSTREAM.send_otp",
	}).ExpectSuccess().Run(t, &response)
	t.Logf("Response TaskID: %s", response.TaskID)

	/* Check in db */
	bucket := response.TaskID[:8]
	taskID := response.TaskID[8:]
	var task task.Task
	if err := qb.Select("scyna.task").
		Columns("bucket", "id", "recurring_task_id", "send_to", "type", "time", "data").
		Where(qb.Eq("bucket"), qb.Eq("id")).
		Query(scyna.DB).
		Bind(bucket, taskID).
		GetRelease(&task); err != nil {
		t.Fatalf("Cannot get task: %s", err.Error())
	}
	t.Logf("Task: %+v\n", task)

	defer qb.Delete("scyna.task").Where(qb.Eq("bucket"), qb.Eq("id")).Query(scyna.DB).Bind(bucket, taskID).ExecRelease()
}
