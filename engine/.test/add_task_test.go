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

func TestAddTask(t *testing.T) {
	var request = scyna.AddTaskRequest{
		Module:   "scheduler",
		Topic:    STREAM_TEST + "." + SUBJECT_TEST,
		Start:    time.Now().Add(time.Second * 10).Unix(),
		Interval: 60,
		LoopMax:  2,
	}
	var response scyna.AddTaskResponse
	scyna_test.ServiceTest(scyna.ADD_TASK_URL).WithRequest(&request).ExpectSuccess().Run(t, &response)
	t.Logf("Response TaskID: %s", response.TaskID)

	/* Check in db */
	taskID, _ := strconv.ParseUint(response.TaskID, 10, 64)
	var task task.Task
	if err := qb.Select("scyna.task").
		Columns("*").
		Where(qb.Eq("id")).
		Query(scyna.DB).
		Bind(taskID).
		GetRelease(&task); err != nil {
		t.Fatalf("Cannot get task: %s", err.Error())
	}
	t.Logf("Task: %+v\n", task)

	bucket := scyna.GetMinuteByTime(time.Unix(request.Start, 0))
	defer qb.Delete("scyna.task").Where(qb.Eq("id")).Query(scyna.DB).Bind(taskID).ExecRelease()
	defer qb.Delete("scyna.todo").Where(qb.Eq("bucket"), qb.Eq("task_id")).Query(scyna.DB).Bind(bucket, taskID).ExecRelease()
}
