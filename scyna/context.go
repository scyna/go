package scyna

import (
	"encoding/json"
	"fmt"
	"log"

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

func (ctx *Context) AuthDone(r proto.Message, token string, expired uint64) {
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
}

// func (s *Context) Auth(org string, secret string, apps []string, userID string) (bool, string) {
// 	request := CreateAuthRequest{
// 		Organization: org,
// 		Secret:       secret,
// 		Apps:         apps,
// 		UserID:       userID,
// 	}

// 	var response CreateAuthResponse
// 	if err := CallService(AUTH_CREATE_URL, &request, &response); err != OK {
// 		return false, ""
// 	}
// 	s.Response.Token = response.Token
// 	s.Response.Expired = response.Expired
// 	return true, response.Token
// }

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
