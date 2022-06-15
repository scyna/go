package task

import "time"

type Task struct {
	Bucket          int64     `db:"bucket"`
	ID              uint64    `db:"id"`
	Time            time.Time `db:"time"`
	RecurringTaskID uint64    `db:"recurring_task_id"`
	SendTo          string    `db:"send_to"`
	Type            string    `db:"type"`
	Data            []byte    `db:"data"`
}

type RecurringTask struct {
	ID       uint64    `db:"id"`
	Time     time.Time `db:"time"`
	Interval int64     `db:"interval"`
	SendTo   string    `db:"send_to"`
	Type     string    `db:"type"`
	Data     []byte    `db:"data"`
	Count    int64     `db:"count"`
}
