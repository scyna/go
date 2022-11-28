package scyna

import (
	"encoding/json"
	"log"
	reflect "reflect"

	"github.com/gocql/gocql"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type CommandHandler[R proto.Message] func(ctx *Command, request R)

func RegisterCommand[R proto.Message](url string, handler CommandHandler[R]) {
	log.Println("Register Command: ", url)
	var request R
	ref := reflect.New(reflect.TypeOf(request).Elem())
	request = ref.Interface().(R)

	ctx := Command{
		Context: Context{Logger{session: false}},
		request: request,
		Batch:   DB.NewBatch(gocql.UnloggedBatch),
	}

	_, err := Connection.QueueSubscribe(SubscriberURL(url), "API", func(m *nats.Msg) {
		if err := proto.Unmarshal(m.Data, &ctx.Request); err != nil {
			log.Print("Register unmarshal error response data:", err.Error())
			return
		}

		ctx.ID = ctx.Request.TraceID
		ctx.Reply = m.Reply
		ctx.Reset(ctx.ID)

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
		log.Fatal("Can not register command:", url)
	}
}
