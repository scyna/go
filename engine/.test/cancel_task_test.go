package test

import (
	"strconv"
	"testing"
	"time"

	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scyna/go/engine/manager/task"
	"github.com/scyna/go/scyna"
	scyna_test "github.com/scyna/go/scyna/testing"
)

func TestCancelTask(t *testing.T) {
	// Add task
	var addRequest = scyna.AddTaskRequest{
		Module:   "scyna",
		Topic:    "TestStream",
		Start:    time.Now().Add(time.Second * 10).Unix(),
		Interval: int64(time.Second * 50),
		LoopMax:  10,
	}
	var addResponse scyna.AddTaskResponse
	scyna_test.ServiceTest(scyna.ADD_TASK_URL).WithRequest(&addRequest).ExpectSuccess().Run(t, &addResponse)
	// Cancel task
	var request = scyna.CancelTaskRequest{
		TaskID: addResponse.TaskID,
	}
	scyna_test.ServiceTest(scyna.CANCEL_TASK_URL).WithRequest(&request).ExpectSuccess().Run(t)
	taskID, _ := strconv.ParseUint(addResponse.TaskID, 10, 64)

	// Check in db
	var task task.Task
	if err := qb.Select("scyna.task").
		Columns("*").
		Where(qb.Eq("id")).
		Query(scyna.DB).
		Bind(task.ID).
		ExecRelease(); err != nil {
		t.Fatalf("Not found task: %s\n", err.Error())
	}

	if task.Active {
		t.Fatalf("Task is active: %+v\n", task)
	}
	defer qb.Delete("scyna.task").Where(qb.Eq("id")).Query(scyna.DB).Bind(taskID).ExecRelease()
}
