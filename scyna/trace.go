package scyna

import (
	"time"

	"google.golang.org/protobuf/proto"
)

type TraceType uint32

const (
	TRACE_SERVICE TraceType = 1
	TRACE_EVENT   TraceType = 2
	TRACE_SIGNAL  TraceType = 3
)

type context struct {
	ParentID uint64
	ID       uint64
	Type     TraceType
	Time     time.Time
	Source   string
}

type Service struct {
	context
	Request Request
	Reply   string
}

func (ctx *context) Tag(key string, value string) {
	/*TODO*/
}

func (ctx *context) Log(level LogLevel, value string) {
	/*TODO*/
}

func (ctx *context) EmitSignal(channel string, data proto.Message) {

	/*TODO*/
}

func (ctx *context) EmitEvent(channel string, data proto.Message) {
	/*TODO*/
}

func (ctx *context) CallService(url string, request proto.Message, response proto.Message) {
	/*TODO*/
}

func (ctx *context) SendCommand(url string, response proto.Message) {
	/*TODO*/
}
