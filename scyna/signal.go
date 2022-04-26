package scyna

import (
	"log"
	reflect "reflect"

	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type SignalStatefulHandler[R proto.Message] func(data R)
type SignalStatelessHandler func()

func RegisterStatefulSignal[R proto.Message](channel string, handler SignalStatefulHandler[R]) {
	var request R
	ref := reflect.New(reflect.TypeOf(request).Elem())
	request = ref.Interface().(R)

	_, err := Connection.QueueSubscribe(channel, module, func(m *nats.Msg) {
		if err := proto.Unmarshal(m.Data, request); err == nil {
			handler(request)
		} else {
			log.Print("Error in parsing data:", err)
		}
	})

	if err != nil {
		log.Fatal("Error in register event")
	}
}

func RegisterStatelessSignal(channel string, handler SignalStatelessHandler) {
	_, err := Connection.QueueSubscribe(channel, module, func(m *nats.Msg) {
		handler()
	})

	if err != nil {
		log.Fatal("Error in register event")
	}
}

func EmitEmptySignal(channel string) {
	var data []byte
	err := Connection.Publish(channel, data)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func EmitSignal(channel string, event proto.Message) {
	data, err := proto.Marshal(event)
	if err != nil {
		log.Print(err.Error())
	}
	if err := Connection.Publish(channel, data); err != nil {
		log.Print(err.Error())
	}
}
