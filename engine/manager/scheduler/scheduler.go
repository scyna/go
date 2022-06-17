package scheduler

import (
	"log"
	"time"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scyna/go/scyna"
)

const INTERVAL = 30 * time.Second
const CLEANUP_FACTOR = 10

var ticker *time.Ticker
var done = make(chan bool)
var cleanupCount = 0

type task struct {
	ID        uint64    `db:"id"`
	Topic     string    `db:"topic"`
	Data      []byte    `db:"data"`
	Active    bool      `db:"active"`
	LoopCount uint64    `db:"loop_count"`
	LoopMax   uint64    `db:"loop_max"`
	Next      time.Time `db:"next"`
	Interval  uint64    `db:"interval"`
}

func Start() {
	ticker = time.NewTicker(INTERVAL)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				loop()
			}
		}
	}()
}

func Stop() {
	ticker.Stop()
	done <- true
}

func loop() {
	bucket := bucket(time.Now())
	for {
		tasks := todos(bucket)
		if tasks == nil {
			break
		}
		for _, t := range tasks {
			process(t)
		}
	}

	cleanupCount++
	if cleanupCount > CLEANUP_FACTOR {
		go cleanup()
		cleanupCount = 0
	}
}

func todos(bucket int64) []int64 {
	var ret []int64
	if err := qb.Select("scyna.todo").
		Columns("id").
		Where(qb.Eq("bucket")).
		Limit(20).
		Query(scyna.DB).
		Bind(bucket).
		SelectRelease(ret); err != nil {
		return nil
	}

	if len(ret) == 0 {
		return nil
	}

	return ret
}

func process(id int64) {
	var t task
	if err := qb.Select("scyna.task").
		Columns("id", "topic", "data", "next", "interval", "loop_count", "loop_max", "running").
		Where(qb.Eq("id")).
		Limit(1).
		Query(scyna.DB).
		Bind(id).
		GetRelease(t); err != nil {
		log.Print("Can not load task")
		return
	}
	if !t.Active {
		/* Do nothing, remove task from todo list only*/
		bucket := scyna.GetMinuteByTime(t.Next)
		if err := qb.Delete("scyna.todo").
			Where(qb.Eq("bucket"), qb.Eq("task_id")).
			Query(scyna.DB).
			Bind(bucket, t.ID).
			ExecRelease(); err != nil {
			scyna.LOG.Error(err.Error())
		}
		return
	}

	// Mark task is doing
	if applied, err := qb.Insert("scyna.doing").
		Columns("task_id").
		Unique().
		TTL(60 * time.Second).
		Query(scyna.DB).
		Bind(t.ID).
		ExecCASRelease(); !applied {
		if err != nil {
			scyna.LOG.Error(err.Error())
		} else {
			scyna.LOG.Error("Task has been doing")
		}
		return
	}

	// SEND signal excute TASK
	scyna.JetStream.Publish(t.Topic, t.Data)

	oldBucket := bucket(t.Next)
	// 1. remove old task on todolist
	qBatch := scyna.DB.NewBatch(gocql.LoggedBatch)
	qBatch.Query("DELETE FROM scyna.todo WHERE bucket = ? AND id = ?;", oldBucket, t.ID)

	t.LoopCount++
	if t.LoopCount < t.LoopMax {
		/* calculate next*/
		t.Next = t.Next.Add(time.Second * time.Duration(t.Interval))
		nextBucket := scyna.GetMinuteByTime(t.Next)
		// 2. add new task to todo list
		qBatch.Query("INSERT INTO scyna.todo (bucket, id) VALUES (?, ?);", nextBucket, t.ID)
	} else {
		t.Active = false
	}
	// 3. update task
	qBatch.Query("UPDATE scyna.task SET next = ?, active = ?, loop_count = ?  WHERE id = ?;", t.Next, t.Active, t.LoopCount, t.ID)
	if err := scyna.DB.ExecuteBatch(qBatch); err != nil {
		scyna.LOG.Error(err.Error())
	}
}

func cleanup() {
	/*TODO*/
}

func bucket(time time.Time) int64 {
	return time.Unix() / 60
}
