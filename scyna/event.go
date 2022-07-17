package scyna

import (
	"log"
	reflect "reflect"
	"time"

	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type EventHandler[R proto.Message] func(ctx *Context, data R)

type eventStream struct {
	sender   string
	receiver string
	handlers map[string]func(m *nats.Msg)
}

var eventStreams map[string]*eventStream = make(map[string]*eventStream)

func RegisterEvent[R proto.Message](sender string, channel string, handler EventHandler[R]) {
	stream := createOrGetEventStream(sender)
	subject := sender + "." + channel
	var event R
	ref := reflect.New(reflect.TypeOf(event).Elem())
	event = ref.Interface().(R)

	trace := Trace{
		Path:      subject,
		SessionID: Session.ID(),
		Type:      TRACE_EVENT,
	}

	stream.handlers[channel] = func(m *nats.Msg) {
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

		if err := proto.Unmarshal(msg.Body, event); err == nil {
			handler(&context, event)
		} else {
			log.Print("Error in parsing data:", err)
		}

		trace.Record()
	}
}

func (es *eventStream) start() {
	sub, err := JetStream.PullSubscribe("", es.receiver, nats.BindStream(es.sender))

	if err != nil {
		log.Fatal("Error in start event stream:", err.Error())
	}

	go func() {
		for {
			if messages, err := sub.Fetch(1); err == nil {
				for _, m := range messages {
					if handler, ok := es.handlers[m.Subject]; ok {
						handler(m)
						m.Ack()
					}
				}
			} else {
				log.Print(err)
			}
		}
	}()
}

func createOrGetEventStream(sender string) *eventStream {
	if stream, ok := eventStreams[sender]; ok {
		return stream
	}

	stream := &eventStream{
		sender:   sender,
		receiver: module,
		handlers: make(map[string]func(m *nats.Msg)),
	}

	eventStreams[sender] = stream
	stream.start()
	return stream
}
