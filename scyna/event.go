package scyna

import (
	"log"
	reflect "reflect"

	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type EventHandler[R proto.Message] func(data R)

func RegisterEvent[R proto.Message](channel string, consumer string, handler EventHandler[R]) {
	var request R
	ref := reflect.New(reflect.TypeOf(request).Elem())
	request = ref.Interface().(R)

	_, err := JetStream.QueueSubscribe(channel, module, func(m *nats.Msg) {
		if err := proto.Unmarshal(m.Data, request); err == nil {
			handler(request)
		} else {
			log.Print("Error in parsing data:", err)
		}

	}, nats.Durable(consumer))

	if err != nil {
		log.Fatal("JetStream Error: ", err)
	}
}

func PostEvent(channel string, event proto.Message) {
	data, err := proto.Marshal(event)
	if err != nil {
		log.Print(err.Error())
	}
	if _, err := JetStream.Publish(channel, data); err != nil {
		log.Print(err.Error())
	}
}
