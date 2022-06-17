package scheduler

import (
	"log"
	"time"

	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scyna/go/scyna"
)

const INTERVAL = 30 * time.Second
const CLEANUP_FACTOR = 10

var ticker *time.Ticker
var done = make(chan bool)
var cleanupCount = 0

type task struct {
	ID        uint64
	Topic     string
	Data      []byte
	Running   bool
	LoopCount uint64
	LoopMax   uint64
	Next      time.Time
	Interval  uint64
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

func todos(bucket uint64) []uint64 {
	var ret []uint64
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

func process(id uint64) {
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
	if !t.Running {
		/*TODO: do nothing, remove task from todo list only*/
		return
	}

	scyna.JetStream.Publish(t.Topic, t.Data)

	t.LoopCount++
	if t.LoopCount < t.LoopMax {
		/*TODO: calculate next*/
		/*Batch:
		1. remove old task on todolist
		2. add new task to todo list
		3. update task
		*/
	} else {
		t.Running = false
		/*Batch:
		1. remove old task on todolist
		3. update task
		*/
	}
}

func cleanup() {
	/*TODO*/
}

func bucket(time time.Time) uint64 {
	return uint64(time.Unix() / 60)
}
