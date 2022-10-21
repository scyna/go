package scyna

import (
	"encoding/json"
	"fmt"
	"log"

	"google.golang.org/protobuf/proto"
)

type Service struct {
	Context
	Request Request
	Reply   string
	request proto.Message
}

func (ctx *Service) Error(e *Error) {
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
	ctx.tag(uint32(response.Code), e)
}

func (ctx *Service) Done(r proto.Message) {
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
	ctx.tag(uint32(response.Code), r)
}

func (ctx *Service) AuthDone(r proto.Message, token string, expired uint64) {
	response := Response{Code: 200, Token: token, Expired: expired}

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
	ctx.tag(200, r)
}

func (ctx *Service) flush(response *Response) {
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

func (ctx *Service) tag(code uint32, response proto.Message) {
	if ctx.ID == 0 {
		return
	}
	res, _ := json.Marshal(response)

	EmitSignal(SERVICE_DONE_CHANNEL, &ServiceDoneSignal{
		TraceID:  ctx.ID,
		Response: string(res),
	})
}
