package scyna

import "time"

type TraceType uint32

const (
	TRACE_SERVICE TraceType = 1
	TRACE_EVENT   TraceType = 2
	TRACE_SIGNAL  TraceType = 3
	TRACE_FUNC    TraceType = 4
)

type Trace struct {
	ParentID uint64
	ID       uint64
	Type     TraceType
	Time     time.Time
	Source   string
}

type Service struct {
	Trace
	Request Request
	Reply   string
}

func (tr *Trace) Save() {
	/*TODO*/
}

func (tr *Trace) Tag(key string, value string) {
	/*TODO*/
}

func (tr *Trace) Log(level LogLevel, value string) {
	/*TODO*/
}

func (tr *Trace) SendRequest() {
	/*TODO*/
}

func (tr *Trace) EmitSignal() {
	/*TODO*/
}

func (tr *Trace) PostEvent() {
	/*TODO*/
}

func (tr *Trace) SendCommand() {
	/*TODO*/
}
