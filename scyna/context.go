package scyna

import (
	"fmt"
	"time"

	"google.golang.org/protobuf/proto"
)

type Context struct {
	Logger
}

func (ctx *Context) Schedule(task string, start time.Time, interval int64, data []byte, loop uint64) (*Error, uint64) {
	var response StartTaskResponse
	if err := ctx.CallService(START_TASK_URL, &StartTaskRequest{
		Module:   Session.context,
		Topic:    fmt.Sprintf("%s.task.%s", Session.context, task),
		Data:     data,
		Time:     start.Unix(),
		Interval: interval,
		Loop:     loop,
	}, &response); err.Code != OK.Code {
		return err, 0
	}

	return nil, response.Id
}

func (ctx *Context) CallService(url string, request proto.Message, response proto.Message) *Error {
	trace := Trace{
		ID:       ID.Next(),
		ParentID: ctx.ID,
		Time:     time.Now(),
		Path:     url,
		Type:     TRACE_SERVICE,
		Source:   Session.context,
	}
	return callService_(&trace, url, request, response)
}

func (ctx *Context) Tag(key string, value string) {
	if ctx.ID == 0 {
		return
	}
	EmitSignal(TAG_CREATED_CHANNEL, &TagCreatedSignal{
		TraceID: ctx.ID,
		Key:     key,
		Value:   value,
	})
}
