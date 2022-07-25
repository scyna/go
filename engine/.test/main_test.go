package test

import (
	"os"
	"testing"
	"time"

	"github.com/scyna/go/engine/manager/authentication"
	"github.com/scyna/go/engine/manager/generator"
	"github.com/scyna/go/engine/manager/logging"
	"github.com/scyna/go/engine/manager/scheduler"
	"github.com/scyna/go/engine/manager/session"
	"github.com/scyna/go/engine/manager/setting"
	"github.com/scyna/go/engine/manager/trace"
	"github.com/scyna/go/scyna"
)

const MODULE_CODE = "scyna.engine"
const STREAM_TEST = "TEST_STREAM"
const SUBJECT_TEST = "SUBJECT"

type task struct {
	ID        uint64    `db:"id"`
	Topic     string    `db:"topic"`
	Data      []byte    `db:"data"`
	Done      bool      `db:"done"`
	LoopIndex uint64    `db:"loop_index"`
	LoopCount uint64    `db:"loop_count"`
	Next      time.Time `db:"next"`
	Start     time.Time `db:"start"`
	Interval  uint64    `db:"interval"`
}

type toDo struct {
	Bucket uint64 `db:"bucket"`
	TaskID uint64 `db:"task_id"`
}

func TestMain(m *testing.M) {
	config := scyna.Configuration{
		NatsUrl:      "localhost",
		NatsUsername: "",
		NatsPassword: "",
		DBHost:       "localhost",
		DBUsername:   "",
		DBPassword:   "",
		DBLocation:   "",
	}

	/* Init module */
	scyna.DirectInit(MODULE_CODE, &config)
	defer scyna.Release()
	generator.Init()
	session.Init(MODULE_CODE, "123456")
	scyna.UseDirectLog(5)

	/* generator */
	scyna.RegisterCommand(scyna.GEN_GET_ID_URL, generator.GetID)
	scyna.RegisterService(scyna.GEN_GET_SN_URL, generator.GetSN)

	/*logging*/
	scyna.RegisterSignalLite(scyna.LOG_CREATED_CHANNEL, logging.Write)

	/*trace*/
	scyna.RegisterSignalLite(scyna.TRACE_CREATED_CHANNEL, trace.TraceCreated)
	scyna.RegisterSignalLite(scyna.TAG_CREATED_CHANNEL, trace.TagCreated)
	scyna.RegisterSignalLite(scyna.SERVICE_DONE_CHANNEL, trace.ServiceDone)

	/*setting*/
	scyna.RegisterService(scyna.SETTING_READ_URL, setting.Read)
	scyna.RegisterService(scyna.SETTING_WRITE_URL, setting.Write)
	scyna.RegisterService(scyna.SETTING_REMOVE_URL, setting.Remove)

	/*authentication*/
	scyna.RegisterService(scyna.AUTH_CREATE_URL, authentication.Create)
	scyna.RegisterService(scyna.AUTH_GET_URL, authentication.Get)
	scyna.RegisterService(scyna.AUTH_LOGOUT_URL, authentication.Logout)

	/* task */
	scyna.RegisterService(scyna.START_TASK_URL, scheduler.StartTask)
	scyna.RegisterService(scyna.STOP_TASK_URL, scheduler.StopTask)

	/* Update config */
	setting.UpdateDefaultConfig(&config)

	/*session*/
	scyna.RegisterSignalLite(scyna.SESSION_END_CHANNEL, session.End)
	scyna.RegisterSignalLite(scyna.SESSION_UPDATE_CHANNEL, session.Update)

	os.Exit(m.Run())
}
