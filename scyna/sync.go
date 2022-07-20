package scyna

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	reflect "reflect"
	"time"

	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type SyncHandler[R proto.Message] func(ctx *Context, data R) *http.Request

func RegisterSync[R proto.Message](channel string, receiver string, handler SyncHandler[R]) {
	subject := module + ".sync." + channel        // vf_account.sync.account
	durable := "sync_" + channel + "_" + receiver // sync_account_loyalty
	LOG.Info(fmt.Sprintf("Channel %s, durable: %s", subject, durable))

	var event R
	ref := reflect.New(reflect.TypeOf(event).Elem())
	event = ref.Interface().(R)

	trace := Trace{
		Path:      subject, //FIXME
		SessionID: Session.ID(),
		Type:      TRACE_SYNC,
	}

	_, err := JetStream.QueueSubscribe(subject, receiver, func(m *nats.Msg) {
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

		if err := proto.Unmarshal(msg.Body, event); err != nil {
			log.Print("Error in parsing data:", err)
			m.Ack()
			return
		}

		request := handler(&context, event)
		if sendSyncRequest(request) {
			m.Ack()
		} else {
			for i := 0; i < 3; i++ {
				request := handler(&context, event)
				if sendSyncRequest(request) {
					m.Ack()
					return
				}
				time.Sleep(time.Second * 30)
			}
			time.Sleep(time.Minute * 10)
			m.Nak()
		}
		trace.Record()
	}, nats.Durable(durable), nats.ManualAck())

	if err != nil {
		log.Fatal("JetStream Error: ", err)
	}
}

func sendSyncRequest(request *http.Request) bool {
	if request == nil {
		return true
	}

	response, err := HttpClient().Do(request)
	if err != nil {
		LOG.Warning("Sync:" + err.Error())
		return false
	} else {
		defer response.Body.Close()
		bodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			LOG.Info("Sync error: " + err.Error())
			return true
		}
		bodyString := string(bodyBytes)
		LOG.Info(fmt.Sprintf("Sync: %s - %d - %s", request.URL, response.StatusCode, bodyString))

		if response.StatusCode == 500 {
			return false
		}
	}
	return true
}
