package scyna

import (
	"fmt"
	"log"
	reflect "reflect"
	"time"

	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type EventHandler[R proto.Message] func(ctx *Context, data R)

type eventStream struct {
	sender    string
	receiver  string
	executors map[string]func(m *nats.Msg)
}

var eventStreams map[string]*eventStream = make(map[string]*eventStream)

func RegisterEvent[R proto.Message](sender string, channel string, handler EventHandler[R]) {
	stream := createOrGetEventStream(sender)
	subject := sender + "." + channel
	LOG.Info(fmt.Sprintf("Events: subject = %s, receiver = %s", subject, stream.receiver))
	var event R
	ref := reflect.New(reflect.TypeOf(event).Elem())
	event = ref.Interface().(R)

	trace := Trace{
		Path:      subject,
		SessionID: Session.ID(),
		Type:      TRACE_EVENT,
	}

	stream.executors[subject] = func(m *nats.Msg) {
		var msg Event
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
			// for _, entityID := range msg.Entities {
			// 	//addActivity(entityID, eventID)
			// }
			// TODO: update entity id to module_name.event_store
		} else {
			log.Print("Error in parsing data:", err)
		}

		trace.Record()
	}
}

func (es *eventStream) start() {
	sub, err := JetStream.PullSubscribe("", es.receiver, nats.BindStream(es.sender))

	if err != nil {
		Fatal("Error in start event stream - sender", es.sender, "- receiver", es.receiver, ":", err.Error())
	}

	go func() {
		for {
			if messages, err := sub.Fetch(1); err == nil {
				if len(messages) == 1 {
					m := messages[0]
					if executor, ok := es.executors[m.Subject]; ok {
						executor(m)
					}
					m.Ack()
				}
			}
		}
	}()
}

func createOrGetEventStream(sender string) *eventStream {
	if stream, ok := eventStreams[sender]; ok {
		return stream
	}

	stream := &eventStream{
		sender:    sender,
		receiver:  context,
		executors: make(map[string]func(m *nats.Msg)),
	}

	eventStreams[sender] = stream
	return stream
}

func startEventStream() {
	for _, e := range eventStreams {
		e.start()
	}
}
