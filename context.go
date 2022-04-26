package scyna

import (
	"encoding/json"
	"log"
	"reflect"

	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type Context struct {
	Request Request
	LOG     Logger
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

		if ctx.Request.LogDisable {
			ctx.LOG = &nullLogger
		} else {
			ctx.LOG = &logger{session: false, ID: ctx.Request.CallID}
		}

		if ctx.Request.JSON {
			if err := json.Unmarshal(ctx.Request.Body, request); err != nil {
				log.Print("Bad Request: " + err.Error())
				//s.Error(BAD_REQUEST)
			} else {
				handler(&ctx, request)
			}

		} else {
			if err := proto.Unmarshal(ctx.Request.Body, request); err != nil {
				log.Print("Bad Request: " + err.Error())
				//s.Error(BAD_REQUEST)
			} else {
				handler(&ctx, request)
			}
		}
	})

	if err != nil {
		log.Fatal("Can not register service:", url)
	}
}

func CreateUser(ctx *Context, request *Request) {

}

func Test() {
	RegisterStatefullService("", CreateUser)
}
