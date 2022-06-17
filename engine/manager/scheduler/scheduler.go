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

var done = make(chan bool)
var cleanupCount = 0

type task struct {
	ID        uint64    `db:"id"`
	Topic     string    `db:"topic"`
	Data      []byte    `db:"data"`
	Done      bool      `db:"done"`
	LoopIndex uint64    `db:"loop_index"`
	LoopCount uint64    `db:"loop_count"`
	Next      time.Time `db:"next"`
	Interval  uint64    `db:"interval"`
}

func Start() {
	ticker := time.NewTicker(INTERVAL)
	go func() {
		for {
			select {
			case <-done:
				ticker.Stop()
				return
			case <-ticker.C:
				loop()
			}
		}
	}()
}

func Stop() {
	done <- true
}

func loop() {
	bucket := getBucket(time.Now())
	for {
		tasks := todos(bucket)
		if tasks == nil {
			break
		}
		for _, task := range tasks {
			if check(bucket, task) {
				process(bucket, task)
			}
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

func check(bucket int64, id int64) bool {
	if applied, _ := qb.Insert("scyna.doing").
		Columns("bucket", "task_id").
		Unique().
		TTL(60*time.Second).
		Query(scyna.DB).
		Bind(bucket, id).
		ExecCASRelease(); !applied {
		return false
	}
	return true
}

func process(bucket int64, id int64) {
	var t task
	if err := qb.Select("scyna.task").
		Columns("id", "topic", "data", "next", "interval", "loop_index", "loop_count", "done").
		Where(qb.Eq("id")).
		Limit(1).
		Query(scyna.DB).
		Bind(id).
		GetRelease(t); err != nil {
		log.Print("Can not load task")
		return
	}

	if bucket != getBucket(t.Next) {
		return /*task is executed somewhere*/
	}

	if t.Done {
		if err := qb.Delete("scyna.todo").
			Where(qb.Eq("bucket"), qb.Eq("task_id")).
			Query(scyna.DB).
			Bind(bucket, id).
			ExecRelease(); err != nil {
			scyna.LOG.Error(err.Error())
		}
		return
	}

	scyna.JetStream.Publish(t.Topic, t.Data) /*activate task handler*/

	qBatch := scyna.DB.NewBatch(gocql.LoggedBatch)
	qBatch.Query("DELETE FROM scyna.todo WHERE bucket = ? AND id = ?;", bucket, id) /* remove old task from todolist */

	t.LoopIndex++
	if t.LoopIndex < t.LoopCount {
		t.Next = t.Next.Add(time.Second * time.Duration(t.Interval)) /* calculate next */
		nextBucket := getBucket(t.Next)
		qBatch.Query("INSERT INTO scyna.todo (bucket, id) VALUES (?, ?);", nextBucket, t.ID) /* add new task to todo list */
		qBatch.Query("UPDATE scyna.task SET next = ?, loop_index = ?  WHERE id = ?;", t.Next, t.LoopIndex, t.ID)
	} else {
		qBatch.Query("UPDATE scyna.task SET done = true WHERE id = ?;")
	}

	if err := scyna.DB.ExecuteBatch(qBatch); err != nil {
		scyna.LOG.Error(err.Error())
	}
}

func cleanup() {
	/*TODO*/
}

func getBucket(time time.Time) int64 {
	return time.Unix() / 60
}
