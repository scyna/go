package scyna

import (
	"encoding/json"
	"log"
	"time"

	"google.golang.org/protobuf/encoding/protojson"

	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type EndpointHandler[R proto.Message] func(ctx *Endpoint, request R) *Error

func RegisterEndpoint[R proto.Message](url string, handler EndpointHandler[R]) {
	log.Println("Register Service: ", url)
	var request R

	ctx := Endpoint{
		Context: Context{Logger{session: false}},
		request: request,
	}

	_, err := Connection.QueueSubscribe(SubscriberURL(url), "API", func(m *nats.Msg) {
		if err := proto.Unmarshal(m.Data, &ctx.Request); err != nil {
			log.Print("Register unmarshal error response data:", err.Error())
			return
		}

		ctx.ID = ctx.Request.TraceID
		ctx.Reply = m.Reply
		ctx.finished = false
		ctx.Reset(ctx.ID)
		ref := request.ProtoReflect().New()
		request = ref.Interface().(R)

		if ctx.Request.JSON {
			if err := json.Unmarshal(ctx.Request.Body, request); err != nil {
				log.Print("Bad Request: " + err.Error())
				ctx.Error(BAD_REQUEST)
			}
		} else {
			if err := proto.Unmarshal(ctx.Request.Body, request); err != nil {
				log.Print("Bad Request: " + err.Error())
				ctx.Error(BAD_REQUEST)
			}
		}

		if !ctx.finished {
			ret := handler(&ctx, request)
			if ret == OK {
				if !ctx.finished {
					ctx.Done(OK)
				}
			} else {
				ctx.Error(ret)
			}
		}
	})

	if err != nil {
		Fatal("Can not register service:", url)
	}
}

func callEndpoint_(trace *Trace, url string, request proto.Message, response proto.Message) *Error {
	defer trace.Record()

	req := Request{TraceID: trace.ID, JSON: false}
	res := Response{}

	if request != nil {
		var err error
		if req.Body, err = proto.Marshal(request); err != nil {
			return BAD_REQUEST
		}
	}

	defer EmitSignalLite(
		TAG_CREATED_CHANNEL,
		&TagCreatedSignal{
			TraceID: trace.ID,
			Key:     "request",
			Value:   protojson.Format(request),
		},
	)

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

	trace.SessionID = res.SessionID
	trace.Status = res.Code
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
