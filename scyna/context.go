package scyna

import (
	"fmt"
	"time"

	"google.golang.org/protobuf/proto"
)

type Context struct {
	Logger
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

func (ctx *Context) PostEvent(channel string, data proto.Message) { // account_created
	subject := module + "." + channel
	msg := EventOrSignal{ParentID: ctx.ID}
	if data, err := proto.Marshal(data); err == nil {
		msg.Body = data
	}

	if data, err := proto.Marshal(&msg); err == nil {
		JetStream.Publish(subject, data)
	}
}

func (ctx *Context) PostEventAndActivity(channel string, data proto.Message, entities []uint64) {
	subject := module + "." + channel
	msg := EventOrSignal{ParentID: ctx.ID, Entities: entities}
	if data, err := proto.Marshal(data); err == nil {
		msg.Body = data
	}

	if data, err := proto.Marshal(&msg); err == nil {
		_, err := JetStream.Publish(subject, data)
		if err != nil {
			ctx.Logger.Error(fmt.Sprintf("Publish to jetstream: %s\n", err.Error()))
		}
	}
}

func (ctx *Context) PostSync(channel string, data proto.Message) { // account_loyalty
	subject := module + ".sync." + channel
	msg := EventOrSignal{ParentID: ctx.ID}
	if data, err := proto.Marshal(data); err == nil {
		msg.Body = data
	}

	if data, err := proto.Marshal(&msg); err == nil {
		JetStream.Publish(subject, data)
	}
}

func (ctx *Context) SendCommand(url string, response proto.Message) *Error {
	return ctx.CallService(url, nil, response)
}

func (ctx *Context) Schedule(task string, start time.Time, interval int64, data []byte, loop uint64) (*Error, uint64) {
	var response StartTaskResponse
	if err := ctx.CallService(START_TASK_URL, &StartTaskRequest{
		Module:   module,
		Topic:    fmt.Sprintf("%s.task.%s", module, task),
		Data:     data,
		Time:     start.Unix(),
		Interval: interval,
		Loop:     loop,
	}, &response); err.Code != OK.Code {
		return err, 0
	}

	return nil, response.Id
}

func (ctx *Context) StopSchedule(taskID uint64) *Error {
	var response Error
	if err := ctx.CallService(STOP_TASK_URL, &StopTaskRequest{
		Id: taskID,
	}, &response); err.Code != OK.Code {
		return err
	}

	return nil
}

func (ctx *Context) CallService(url string, request proto.Message, response proto.Message) *Error {
	trace := Trace{
		ID:       ID.Next(),
		ParentID: ctx.ID,
		Time:     time.Now(),
		Path:     url,
		Type:     TRACE_SERVICE,
		Source:   module,
	}
	return callService_(&trace, url, request, response)
}

func (ctx *Context) CallEndpoint(url string, request proto.Message, response proto.Message) *Error {
	trace := Trace{
		ID:       ID.Next(),
		ParentID: ctx.ID,
		Time:     time.Now(),
		Path:     url,
		Type:     TRACE_SERVICE,
		Source:   module,
	}
	return callEndpoint_(&trace, url, request, response)
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
