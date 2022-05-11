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

	context := Context{
		Path:      channel,
		SessionID: Session.ID(),
		Type:      TRACE_SIGNAL,
		LOG:       &logger{session: false},
	}

	if _, err := JetStream.QueueSubscribe(channel, module, func(m *nats.Msg) {
		var msg EventOrSignal
		if err := proto.Unmarshal(m.Data, &msg); err != nil {
			log.Print("Register unmarshal error response data:", err.Error())
			return
		}
		context.Time = time.Now()
		context.ID = ID.Next()
		context.ParentID = msg.ParentID
		context.LOG.Reset(context.ID)

		if err := proto.Unmarshal(m.Data, event); err == nil {
			handler(&context, event)
		} else {
			log.Print("Error in parsing data:", err)
		}

	}, nats.Durable(consumer)); err != nil {
		log.Fatal("Error in registering Event: ", err)
	}
}
