package scheduler

import (
	"errors"
	"time"

	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scyna/go/scyna"
)

type Task struct {
	ID        uint64    `db:"id"`
	Topic     string    `db:"topic"`
	Data      []byte    `db:"data"`
	Start     time.Time `db:"start"`
	Next      time.Time `db:"next"`
	Interval  uint64    `db:"interval"`
	LoopCount uint64    `db:"loop_count"`
	LoopIndex uint64    `db:"loop_index"`
	Done      bool      `db:"done"`
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
	Bucket int64  `db:"bucket"`
	TaskID uint64 `db:"task_id"`
}

func (task *Task) Get() error {
	if err := qb.Select("scyna.task").
		Columns("id", "topic", "data", "next", "interval", "loop_count", "loop_max", "done").
		Where(qb.Eq("id")).
		Limit(1).
		Query(scyna.DB).
		Bind(task.ID).
		GetRelease(task); err != nil {
		scyna.LOG.Error(err.Error())
		return err
	}
	return nil
}

func (task *Task) Acquire() error {
	// Mark task is doing
	bucket := getBucket(task.Next)
	if applied, err := qb.Insert("scyna.doing").
		Columns("bucket", "task_id").
		Unique().
		TTL(60*time.Second).
		Query(scyna.DB).
		Bind(bucket, task.ID).
		ExecCASRelease(); !applied {
		if err == nil {
			err = errors.New("Task has been doing")
		}
		scyna.LOG.Error(err.Error())
		return err
	}
	return nil
}

func (task *Task) Deactive() error {
	// Mark task is doing
	if err := qb.Update("scyna.task").
		Set("done").
		Where(qb.Eq("id")).
		Query(scyna.DB).
		Bind(false, task.ID).
		ExecRelease(); err != nil {
		return err
	}
	return nil
}
