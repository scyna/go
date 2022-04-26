package scyna

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type Context struct {
	Request Request
	Reply   string
	LOG     Logger
}

func (ctx *Context) Error(e *Error) {
	response := Response{Code: 400}

	var err error
	if ctx.Request.JSON {
		response.Body, err = json.Marshal(e)
	} else {
		response.Body, err = proto.Marshal(e)
	}

	if err != nil {
		response.Code = int32(500)
		response.Body = []byte(err.Error())
	}
	ctx.flush(&response)
}

func (ctx *Context) Done(r proto.Message) {
	response := Response{Code: 200}

	var err error
	if ctx.Request.JSON {
		response.Body, err = json.Marshal(r)
	} else {
		response.Body, err = proto.Marshal(r)
	}
	if err != nil {
		response.Code = int32(500)
		response.Body = []byte(err.Error())
	}

	ctx.flush(&response)
}

func (ctx *Context) flush(response *Response) {
	response.SessionID = Session.ID()
	bytes, err := proto.Marshal(response)
	if err != nil {
		log.Print("Register marshal error response data:", err.Error())
		return
	}
	err = Connection.Publish(ctx.Reply, bytes)
	if err != nil {
		LOG.Error(fmt.Sprintf("Nats publish to [%s] error: %s", ctx.Reply, err.Error()))
	}
}

type StatefulServiceHandler[R proto.Message] func(ctx *Context, request R)
type StatelessServiceHandler func(ctx *Context)

func RegisterStatefullService[R proto.Message](url string, handler StatefulServiceHandler[R]) {
	log.Println("[Register] Sub url: ", url)
	var request R
	ref := reflect.New(reflect.TypeOf(request).Elem())
	request = ref.Elem().Interface().(R)

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
