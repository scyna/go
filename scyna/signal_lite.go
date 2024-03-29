package scyna

import (
	"log"
	reflect "reflect"

	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type SignalLiteHandler[R proto.Message] func(data R)

func RegisterSignalLite[R proto.Message](channel string, handler SignalLiteHandler[R]) {
	log.Print("Register SignalLite:", channel)
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
		Fatal("Error in register SignalLite")
	}
}

func EmitSignalLite(channel string, event proto.Message) {
	if data, err := proto.Marshal(event); err == nil {
		Connection.Publish(channel, data)
	}
}
