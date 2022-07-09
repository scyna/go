package scyna

import (
	"log"
	reflect "reflect"
	"time"

	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type EventHandler[R proto.Message] func(ctx *Context, data R)

func RegisterEvent[R proto.Message](sender string, channel string, handler EventHandler[R]) {
	consumer := GetEventConsumer(sender, channel, module)
	subject := GetEventSubject(sender, channel)
	var event R
	ref := reflect.New(reflect.TypeOf(event).Elem())
	event = ref.Interface().(R)

	trace := Trace{
		Path:      subject,
		SessionID: Session.ID(),
		Type:      TRACE_EVENT,
	}

	if _, err := JetStream.QueueSubscribe(subject, module, func(m *nats.Msg) {
		var msg EventOrSignal
		defer m.Ack() //assure ordering
		if err := proto.Unmarshal(m.Data, &msg); err != nil {
			log.Print("Register unmarshal error response data:", err.Error())
			return
		}
		trace.Time = time.Now()
		trace.ID = ID.Next()
		trace.ParentID = msg.ParentID

		context := Context{
			Logger{ID: trace.ID, session: false},
		}

		if err := proto.Unmarshal(msg.Body, event); err == nil {
			handler(&context, event)
		} else {
			log.Print("Error in parsing data:", err)
		}

		trace.Record()
	}, nats.Durable(consumer), nats.ManualAck()); err != nil {
		log.Fatal("Error in registering Event: ", err)
	}
}
