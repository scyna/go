package scyna

import (
	"encoding/json"
	"log"
	reflect "reflect"
	"time"

	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type StatefulServiceHandler[R proto.Message] func(ctx *Context, request R)
type StatelessServiceHandler func(ctx *Context)

var nullLogger loggerNull

func CallService(url string, request proto.Message, response proto.Message) *Error {
	req := Request{CallID: ID.Next(), JSON: false}
	res := Response{}

	if request != nil {
		var err error
		if req.Body, err = proto.Marshal(request); err != nil {
			return BAD_REQUEST
		}
	}

	if data, err := proto.Marshal(&req); err == nil {
		if msg, err := Connection.Request(PublishURL(url), data, 10*time.Second); err == nil {
			if err := proto.Unmarshal(msg.Data, &res); err != nil {
				return SERVER_ERROR
			}
		} else {
			return SERVER_ERROR
		}
	} else {
		return BAD_REQUEST
	}

	if res.Code == 200 {
		if err := proto.Unmarshal(res.Body, response); err == nil {
			return OK
		}
	} else {
		var ret Error
		if err := proto.Unmarshal(res.Body, &ret); err == nil {
			return &ret
		}
	}
	return SERVER_ERROR
}

func RegisterStatefulService[R proto.Message](url string, handler StatefulServiceHandler[R]) {
	log.Println("Register Service: ", url)
	var request R
	ref := reflect.New(reflect.TypeOf(request).Elem())
	request = ref.Interface().(R)

	var ctx Context
	_, err := Connection.QueueSubscribe(SubscribreURL(url), "API", func(m *nats.Msg) {
		if err := proto.Unmarshal(m.Data, &ctx.Request); err != nil {
			log.Print("Register unmarshal error response data:", err.Error())
			return
		}

		ctx.Reply = m.Reply
		if ctx.Request.LogDisable {
			ctx.LOG = &nullLogger
		} else {
			ctx.LOG = &logger{session: false, ID: ctx.Request.CallID}
		}

		if ctx.Request.JSON {
			if err := json.Unmarshal(ctx.Request.Body, request); err != nil {
				log.Print("Bad Request: " + err.Error())
				ctx.Error(BAD_REQUEST)
			} else {
				handler(&ctx, request)
			}

		} else {
			if err := proto.Unmarshal(ctx.Request.Body, request); err != nil {
				log.Print("Bad Request: " + err.Error())
				ctx.Error(BAD_REQUEST)
			} else {
				handler(&ctx, request)
			}
		}
	})

	if err != nil {
		log.Fatal("Can not register service:", url)
	}
}

func RegisterStatelessService(url string, handler StatelessServiceHandler) {
	log.Println("[Register] Sub url: ", url)
	var ctx Context
	_, err := Connection.QueueSubscribe(SubscribreURL(url), "API", func(m *nats.Msg) {
		if err := proto.Unmarshal(m.Data, &ctx.Request); err != nil {
			log.Print("Register unmarshal error response data:", err.Error())
			return
		}

		ctx.Reply = m.Reply
		if ctx.Request.LogDisable {
			ctx.LOG = &nullLogger
		} else {
			ctx.LOG = &logger{session: false, ID: ctx.Request.CallID}
		}
		handler(&ctx)
	})

	if err != nil {
		log.Fatal("Can not register service:", url)
	}
}
