package scyna

import (
	"log"
	reflect "reflect"

	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type SignalHandler[R proto.Message] func(LOG Logger, data R)

func RegisterSignal[R proto.Message](channel string, handler SignalHandler[R]) {
	log.Print("Register signal:", channel)
	var signal R
	ref := reflect.New(reflect.TypeOf(signal).Elem())
	signal = ref.Interface().(R)
	LOG := &logger{session: false}

	if _, err := Connection.QueueSubscribe(channel, module, func(m *nats.Msg) {
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
	}); err != nil {
		log.Fatal("Error in register event")
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
