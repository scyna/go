package scyna

import (
	"log"

	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type CommandHandler func(ctx *Service)

func RegisterCommand(url string, handler CommandHandler) {
	log.Println("Register Command:", url)
	ctx := Service{
		Context: Context{Logger{session: false}},
	}

	_, err := Connection.QueueSubscribe(SubscriberURL(url), "API", func(m *nats.Msg) {
		if err := proto.Unmarshal(m.Data, &ctx.Request); err != nil {
			log.Print("Register unmarshal error response data:", err.Error())
			return
		}

		ctx.ID = ctx.Request.TraceID
		ctx.Reply = m.Reply
		ctx.Reset(ctx.ID)
		handler(&ctx)
	})

	if err != nil {
		Fatal("Can not register command:", url)
	}
}
