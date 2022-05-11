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

type Context struct {
	ParentID  uint64    `db:"parent_id"`
	ID        uint64    `db:"id"`
	Type      TraceType `db:"type"`
	Time      time.Time `db:"time"`
	Duration  uint64    `db:"duration"`
	Path      string    `db:"path"`
	Source    string    `db:"source"`
	SessionID uint64    `db:"session_id"`
	Status    int32
	LOG       Logger
}

func (ctx *Context) EmitSignal(channel string, data proto.Message) {
	msg := EventOrSignal{ParentID: ctx.ID}
	if data, err := proto.Marshal(data); err == nil {
		msg.Body = data
	}

	if data, err := proto.Marshal(&msg); err == nil {
		Connection.Publish(channel, data)
	}
}

func (ctx *Context) PostEvent(channel string, data proto.Message) {
	msg := EventOrSignal{ParentID: ctx.ID}
	if data, err := proto.Marshal(data); err == nil {
		msg.Body = data
	}

	if data, err := proto.Marshal(&msg); err == nil {
		JetStream.Publish(channel, data)
	}
}

func (ctx *Context) SendCommand(url string, response proto.Message) *Error {
	return ctx.CallService(url, nil, response)
}

func (ctx *Context) CallService(url string, request proto.Message, response proto.Message) *Error {
	context := Context{
		ID:       ID.Next(),
		ParentID: ctx.ID,
		Time:     time.Now(),
		Path:     url,
		Type:     TRACE_SERVICE,
		Source:   module,
	}
	return callService_(&context, url, request, response)
}

func (ctx *Context) Tag(key string, value string) {
	if ctx.ID == 0 {
		return
	}
	EmitSignalLite(TAG_CREATED_CHANNEL, &TagCreatedSignal{
		TraceID: ctx.ID,
		Key:     key,
		Value:   value,
	})
}

func (ctx *Context) Save() {
	ctx.Duration = uint64(time.Now().UnixNano() - ctx.Time.UnixNano())
	EmitSignalLite(TRACE_CREATED_CHANNEL, &TraceCreatedSignal{
		ID:        ctx.ID,
		ParentID:  ctx.ParentID,
		Type:      uint32(ctx.Type),
		Time:      uint64(ctx.Time.UnixMicro()),
		Duration:  ctx.Duration,
		Path:      ctx.Path,
		Source:    ctx.Source,
		SessionID: ctx.SessionID,
		Status:    ctx.Status,
	})
}
