package scyna

import (
	"log"
	reflect "reflect"
	"time"

	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type EventHandler[R proto.Message] func(ctx *Context, data R)

func RegisterEvent[R proto.Message](channel string, consumer string, handler EventHandler[R]) {
	var event R
	ref := reflect.New(reflect.TypeOf(event).Elem())
	event = ref.Interface().(R)

	trace := Trace{
		Path:      channel,
		SessionID: Session.ID(),
		Type:      TRACE_SIGNAL,
	}

	if _, err := JetStream.QueueSubscribe(channel, module, func(m *nats.Msg) {
		var msg EventOrSignal
		if err := proto.Unmarshal(m.Data, &msg); err != nil {
			log.Print("Register unmarshal error response data:", err.Error())
			return
		}
		trace.Time = time.Now()
		trace.ID = ID.Next()
		trace.ParentID = msg.ParentID

		context := Context{
			ID:  trace.ID,
			LOG: &logger{ID: trace.ID, session: false},
		}

		if err := proto.Unmarshal(m.Data, event); err == nil {
			handler(&context, event)
		} else {
			log.Print("Error in parsing data:", err)
		}

		trace.Save()
	}, nats.Durable(consumer)); err != nil {
		log.Fatal("Error in registering Event: ", err)
	}
}
