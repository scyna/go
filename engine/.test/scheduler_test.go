package test

import (
	"testing"
	"time"

	"github.com/scyna/go/scyna"
	scyna_test "github.com/scyna/go/scyna/testing"
)

func TestProcessInScheduler(t *testing.T) {
	var request = scyna.AddTaskRequest{
		Module:   "scheduler",
		Topic:    STREAM_TEST + "." + SUBJECT_TEST,
		Start:    time.Now().Add(time.Second * 10).Unix(),
		Interval: 60,
		LoopMax:  2,
	}
	var response scyna.AddTaskResponse
	scyna_test.ServiceTest(scyna.ADD_TASK_URL).WithRequest(&request).ExpectSuccess().Run(t, &response)

}
