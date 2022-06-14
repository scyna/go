package task

type Task struct {
	Period int64  `db:"period"`
	Time   int64  `db:"time"`
	Repeat bool   `db:"repeat"`
	SendTo string `db:"send_to"`
	Type   string `db:"type"`
	Data   []byte `db:"data"`
	// TODO: fill fields of scyna.task table
}
