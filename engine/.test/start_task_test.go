package test

import (
	"testing"
	"time"

	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scyna/go/engine/manager/scheduler"
	"github.com/scyna/go/scyna"
	scyna_test "github.com/scyna/go/scyna/testing"
)

func TestStartTask(t *testing.T) {
	var request = scyna.StartTaskRequest{
		Module:   "scheduler",
		Topic:    STREAM_TEST + "." + SUBJECT_TEST,
		Time:     uint64(time.Now().Add(time.Second * 10).Unix()),
		Interval: 60,
		Loop:     2,
	}
	var response scyna.StartTaskResponse
	scyna_test.ServiceTest(scyna.START_TASK_URL).WithRequest(&request).ExpectSuccess().Run(t, &response)
	t.Logf("Response TaskID: %d", response.Id)

	bucket := scyna.GetMinuteByTime(time.Unix(int64(request.Time), 0))
	defer qb.Delete("scyna.todo").Where(qb.Eq("bucket"), qb.Eq("task_id")).Query(scyna.DB).Bind(bucket, response.Id).ExecRelease()
	defer qb.Delete("scyna.task").Where(qb.Eq("id")).Query(scyna.DB).Bind(response.Id).ExecRelease()
	/* Check in db */
	var task scheduler.Task
	if err := qb.Select("scyna.task").
		Columns("*").
		Where(qb.Eq("id")).
		Query(scyna.DB).
		Bind(response.Id).
		GetRelease(&task); err != nil {
		t.Fatalf("Cannot get task: %s", err.Error())
	}
	t.Logf("Task: %+v\n", task)
	if task.Start != task.Next {
		t.Fatalf("Wrong value in time")
	}

	var todo scheduler.ToDo
	if err := qb.Select("scyna.todo").
		Columns("*").
		Where(qb.Eq("bucket"), qb.Eq("task_id")).
		Query(scyna.DB).
		Bind(bucket, response.Id).
		GetRelease(&todo); err != nil {
		t.Fatalf("Cannot get task: %s", err.Error())
	}
	t.Logf("Todo: %+v\n", todo)
}
