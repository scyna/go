package test

import (
	"testing"
	"time"

	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scyna/go/engine/manager/scheduler"
	"github.com/scyna/go/scyna"
	scyna_test "github.com/scyna/go/scyna/testing"
)

func TestStopTask(t *testing.T) {
	// Add task
	var addRequest = scyna.StartTaskRequest{
		Module:   "scyna",
		Topic:    "TestStream",
		Time:     uint64(time.Now().Add(time.Second * 10).Unix()),
		Interval: uint64(time.Second * 50),
		Loop:     10,
	}
	var addResponse scyna.StartTaskResponse
	scyna_test.ServiceTest(scyna.START_TASK_URL).WithRequest(&addRequest).ExpectSuccess().Run(t, &addResponse)
	t.Logf("TaskID: %d\n", addResponse.Id)
	// Cancel task
	var request = scyna.StopTaskRequest{
		Id: addResponse.Id,
	}
	scyna_test.ServiceTest(scyna.STOP_TASK_URL).WithRequest(&request).ExpectSuccess().Run(t)

	// Check in db
	var task scheduler.Task
	task.ID = addResponse.Id
	if err := qb.Select("scyna.task").
		Columns("*").
		Where(qb.Eq("id")).
		Query(scyna.DB).
		Bind(task.ID).
		GetRelease(&task); err != nil {
		t.Fatalf("Not found task: %s\n", err.Error())
	}

	t.Logf("Task: %+v", task)
	if task.Done {
		t.Fatalf("Task is active: %+v\n", task)
	}
	bucket := task.Next.Unix() / 60
	defer qb.Delete("scyna.task").Where(qb.Eq("id")).Query(scyna.DB).Bind(addResponse.Id).ExecRelease()
	if err := qb.Delete("scyna.todo").Where(qb.Eq("bucket"), qb.Eq("task_id")).Query(scyna.DB).Bind(bucket, addResponse.Id).ExecRelease(); err != nil {
		t.Errorf(err.Error())
	}
}
