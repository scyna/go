package task

import (
	"time"
)

type Task struct {
	ID        uint64    `db:"id"`
	Topic     string    `db:"topic"`
	Data      []byte    `db:"data"`
	Start     time.Time `db:"start"`
	Next      time.Time `db:"next"`
	Interval  int64     `db:"interval"`
	LoopCount int64     `db:"loop_count"`
	LoopMax   int64     `db:"loop_max"`
	Active    bool      `db:"active"`
}

type ModuleHasTask struct {
	Module string `db:"module"`
	TaskID uint64 `db:"task_id"`
}

type ToDo struct {
	Bucket int64  `db:"bucket"`
	TaskID uint64 `db:"task_id"`
}

type Doing struct {
	TaskID uint64 `db:"task_id"`
}

func (task *Task) NearestTime() int64 {
	now := time.Now().Unix()
	nInterval := (now-task.Start.Unix())/task.Interval + 1
	return nInterval*task.Interval + task.Start.Unix()
}
