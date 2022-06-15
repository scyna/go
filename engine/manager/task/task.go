package task

type Task struct {
	Period int64  `db:"period"`
	Time   int64  `db:"time"`
	Repeat bool   `db:"repeat"`
	SendTo string `db:"send_to"`
	Type   string `db:"type"`
	Data   []byte `db:"data"`
}

type RecurringTask struct {
	Time     int64  `db:"time"`
	Interval int64  `db:"interval"`
	SendTo   string `db:"send_to"`
	Type     string `db:"type"`
	Data     []byte `db:"data"`
}
