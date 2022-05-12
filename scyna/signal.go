package scyna

import (
	"log"
	reflect "reflect"
	"time"

	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type SignalHandler[R proto.Message] func(ctx *Context, data R)

func RegisterSignal[R proto.Message](channel string, handler SignalHandler[R]) {
	log.Print("Register Signal:", channel)
	var signal R
	ref := reflect.New(reflect.TypeOf(signal).Elem())
	signal = ref.Interface().(R)

	trace := Trace{
		Path:      channel,
		SessionID: Session.ID(),
		Type:      TRACE_SIGNAL,
	}

	if _, err := Connection.QueueSubscribe(channel, module, func(m *nats.Msg) {
		var msg EventOrSignal
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

		if err := proto.Unmarshal(msg.Body, signal); err == nil {
			handler(&context, signal)
		} else {
			log.Print("Error in parsing data:", err)
		}

		trace.Save()
	}); err != nil {
		log.Fatal("Error in register Signal:", err)
	}
}
