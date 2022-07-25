package scheduler

import (
	"time"
)

const INTERVAL = 60 * time.Second

var done = make(chan bool)
var w *worker //you can declare multi worker here

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

func Start(interval time.Duration) {
	w = NewWorker()
	w.Start(0, interval) //FIXME: random
}

func Stop() {
	done <- true
}

func GetBucket(time time.Time) int64 {
	return time.Unix() / 60
}
