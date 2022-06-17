package scheduler

import "time"

const INTERVAL = 30 * time.Second
const CLEANUP_FACTOR = 10

var ticker *time.Ticker
var done = make(chan bool)
var cleanupCount = 0

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
	var bucket uint64
	for {
		tasks := todos(bucket)
		if len(tasks) == 0 {
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
	/*TODO*/
	return nil
}

func process(task uint64) {
	/*TODO*/
}

func cleanup() {
	/*TODO*/
}
