package test

import (
	"testing"
	"time"

	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scyna/go/engine/manager/scheduler"
	"github.com/scyna/go/scyna"
	scyna_test "github.com/scyna/go/scyna/testing"
)

func TestProcessInScheduler(t *testing.T) {
	taskTime := time.Now().Add(time.Second * 2)
	var request = scyna.StartTaskRequest{
		Module:   "scheduler",
		Topic:    STREAM_TEST + "." + SUBJECT_TEST,
		Time:     uint64(taskTime.Unix()),
		Interval: 60,
		Loop:     2,
	}

	var response scyna.StartTaskResponse
	scyna_test.ServiceTest(scyna.START_TASK_URL).WithRequest(&request).ExpectSuccess().Run(t, &response)

	nextBucket := scheduler.GetBucket(taskTime.Add(60 * time.Second))
	defer qb.Delete("scyna.task").Where(qb.Eq("id")).Query(scyna.DB).Bind(response.Id).ExecRelease()
	defer qb.Delete("scyna.todo").Where(qb.Eq("bucket"), qb.Eq("task_id")).Query(scyna.DB).Bind(nextBucket, response.Id).ExecRelease()
	defer qb.Delete("scyna.doing").Where(qb.Eq("bucket"), qb.Eq("task_id")).Query(scyna.DB).Bind(nextBucket-1, response.Id).ExecRelease()

	time.Sleep(time.Second * 1)
	scheduler.Start(time.Second * 2)
	time.Sleep(time.Second * 3)

	var task task
	if err := qb.Select("scyna.task").
		Columns("*").
		Where(qb.Eq("id")).
		Query(scyna.DB).
		Bind(response.Id).
		GetRelease(&task); err != nil {
		t.Fatalf("Cannot get task: %s\n", err.Error())
	}
	t.Logf("Task: %+v\n", task)
	if task.LoopIndex == 0 {
		t.Fatalf("Task is not processed")
	}

	if !task.Done && task.Start == task.Next {
		t.Fatalf("Do not calculate next time!")
	}
}
