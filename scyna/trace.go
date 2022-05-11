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
	defer context.Save()

	req := Request{TraceID: context.ID, JSON: false}
	res := Response{}

	if request != nil {
		var err error
		if req.Body, err = proto.Marshal(request); err != nil {
			return BAD_REQUEST
		}
	}

	if data, err := proto.Marshal(&req); err == nil {
		if msg, err := Connection.Request(PublishURL(url), data, 10*time.Second); err == nil {
			if err := proto.Unmarshal(msg.Data, &res); err != nil {
				return SERVER_ERROR
			}
		} else {
			return SERVER_ERROR
		}
	} else {
		return BAD_REQUEST
	}

	context.SessionID = res.SessionID
	context.Status = res.Code
	if res.Code == 200 {
		if err := proto.Unmarshal(res.Body, response); err == nil {
			return OK
		}
	} else {
		var ret Error
		if err := proto.Unmarshal(res.Body, &ret); err == nil {
			return &ret
		}
	}
	return SERVER_ERROR
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
	})
}
