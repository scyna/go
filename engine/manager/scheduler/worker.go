package scheduler

import (
	"log"
	"time"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scyna/go/scyna"
)

type worker struct {
	qCheck *gocqlx.Queryx
	qGet   *qb.SelectBuilder
	qTodos *gocqlx.Queryx
}

func NewWorker() *worker {
	return &worker{
		qCheck: qb.Insert("scyna.doing").
			Columns("bucket", "task_id").
			Unique().
			TTL(60 * time.Second).
			Query(scyna.DB),
		qGet: qb.Select("scyna.task").
			Columns("id", "topic", "data", "next", "interval", "loop_index", "loop_count", "done").
			Where(qb.Eq("id")).
			Limit(1),
		qTodos: qb.Select("scyna.todo").
			Columns("task_id").
			Where(qb.Eq("bucket")).
			Limit(20).
			Query(scyna.DB),
	}
}

func (w *worker) Start(delay time.Duration, interval time.Duration) {
	go func() {
		time.Sleep(delay)
		ticker := time.NewTicker(interval)
		for {
			select {
			case <-done:
				ticker.Stop()
				return
			case <-ticker.C:
				w.execute()
			}
		}
	}()
}

func (w *worker) execute() {
	bucket := GetBucket(time.Now())
	for {
		var tasks []int64
		if err := w.qTodos.Bind(bucket).Select(&tasks); err != nil || len(tasks) == 0 {
			break
		}

		for _, task := range tasks {
			if applied, _ := w.qCheck.Bind(bucket, task).ExecCAS(); applied {
				w.process(bucket, task)
			}
		}
	}
}

func (w *worker) process(bucket int64, id int64) {
	var t task
	if err := w.qGet.Query(scyna.DB).Bind(id).GetRelease(&t); err != nil {
		log.Print("Can not load task")
		return
	}
	if bucket != GetBucket(t.Next) {
		return
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
	qBatch.Query("DELETE FROM scyna.todo WHERE bucket = ? AND task_id = ?;", bucket, id) /* remove old task from todolist */

	t.LoopIndex++
	if t.LoopIndex < t.LoopCount {
		t.Next = t.Next.Add(time.Second * time.Duration(t.Interval)) /* calculate next */
		nextBucket := GetBucket(t.Next)
		qBatch.Query("INSERT INTO scyna.todo (bucket, task_id) VALUES (?, ?);", nextBucket, t.ID) /* add new task to todo list */
		qBatch.Query("UPDATE scyna.task SET next = ?, loop_index = ?  WHERE id = ?;", t.Next, t.LoopIndex, t.ID)
	} else {
		qBatch.Query("UPDATE scyna.task SET done = true WHERE id = ?;", t.ID)
	}

	if err := scyna.DB.ExecuteBatch(qBatch); err != nil {
		scyna.LOG.Error(err.Error())
	}
}
