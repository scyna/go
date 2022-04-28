package scyna

import (
	"log"
	reflect "reflect"

	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type SignalStatefulHandler[R proto.Message] func(LOG Logger, data R)
type SignalStatelessHandler func(LOG Logger)

func RegisterStatefulSignal[R proto.Message](channel string, handler SignalStatefulHandler[R]) {
	var signal R
	ref := reflect.New(reflect.TypeOf(signal).Elem())
	signal = ref.Interface().(R)
	LOG := &logger{session: false}

	_, err := Connection.QueueSubscribe(channel, module, func(m *nats.Msg) {
		var msg EventOrSignal
		if err := proto.Unmarshal(m.Data, &msg); err != nil {
			log.Print("Register unmarshal error response data:", err.Error())
			return
		}
		LOG.Reset(msg.CallID)

		if err := proto.Unmarshal(msg.Body, signal); err == nil {
			handler(LOG, signal)
		} else {
			log.Print("Error in parsing data:", err)
		}
	})

	if err != nil {
		log.Fatal("Error in register event")
	}
}

func RegisterStatelessSignal(channel string, handler SignalStatelessHandler) {
	LOG := &logger{session: false}
	_, err := Connection.QueueSubscribe(channel, module, func(m *nats.Msg) {
		var msg EventOrSignal
		if err := proto.Unmarshal(m.Data, &msg); err != nil {
			log.Print("Register unmarshal error response data:", err.Error())
			return
		}
		LOG.Reset(msg.CallID)
		handler(LOG)
	})

	if err != nil {
		log.Fatal("Error in register event")
	}
}

func EmitStatelessSignal(channel string) {
	msg := EventOrSignal{CallID: ID.Next()}

	if data, err := proto.Marshal(&msg); err == nil {
		Connection.Publish(channel, data)
	}
}

func EmitSignal(channel string, event proto.Message) {
	msg := EventOrSignal{CallID: ID.Next()}
	if data, err := proto.Marshal(event); err == nil {
		msg.Body = data
	}

	if data, err := proto.Marshal(&msg); err == nil {
		Connection.Publish(channel, data)
	}
}
