package scyna

import (
	"log"

	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type ServiceLiteHandler func(ctx *Context)

func RegisterServiceLite(url string, handler ServiceLiteHandler) {
	log.Println("[Register] Sub url: ", url)
	ctx := Context{LOG: &logger{session: false}}
	_, err := Connection.QueueSubscribe(SubscriberURL(url), "API", func(m *nats.Msg) {
		if err := proto.Unmarshal(m.Data, &ctx.Request); err != nil {
			log.Print("Register unmarshal error response data:", err.Error())
			return
		}

		ctx.Reply = m.Reply
		ctx.LOG.Reset(ctx.Request.CallID)
		handler(&ctx)
	})

	if err != nil {
		log.Fatal("Can not register service:", url)
	}
}
