package scyna

import (
	"log"
	reflect "reflect"

	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type CommandHandler[R proto.Message] func(data R)
type SignalLiteHandler func(LOG Logger)

func RegisterCommand[R proto.Message](channel string, handler CommandHandler[R]) {
	log.Print("Register signal:", channel)
	var signal R
	ref := reflect.New(reflect.TypeOf(signal).Elem())
	signal = ref.Interface().(R)

	if _, err := Connection.QueueSubscribe(channel, module, func(m *nats.Msg) {
		if err := proto.Unmarshal(m.Data, signal); err == nil {
			handler(signal)
		} else {
			log.Print("Error in parsing data:", err)
		}
	}); err != nil {
		log.Fatal("Error in register event")
	}
}

func RegisterSignalLite(channel string, handler SignalLiteHandler) {
	log.Print("Register signal:", channel)
	LOG := &logger{session: false}
	if _, err := Connection.QueueSubscribe(channel, module, func(m *nats.Msg) {
		var msg EventOrSignal
		if err := proto.Unmarshal(m.Data, &msg); err != nil {
			log.Print("Register unmarshal error response data:", err.Error())
			return
		}
		LOG.Reset(msg.CallID)
		handler(LOG)
	}); err != nil {
		log.Fatal("Error in register event")
	}
}

func EmitSignalLite(channel string) {
	msg := EventOrSignal{CallID: ID.Next()}

	if data, err := proto.Marshal(&msg); err == nil {
		Connection.Publish(channel, data)
	}
}

func SendCommand(channel string, event proto.Message) {
	if data, err := proto.Marshal(event); err == nil {
		Connection.Publish(channel, data)
	}
}
