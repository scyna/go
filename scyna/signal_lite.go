package scyna

import (
	"log"

	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type SignalLiteHandler func(LOG Logger)

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
