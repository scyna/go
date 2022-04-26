package scyna

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"sync"
)

type HttpContext struct {
	Request  Request
	Response Response
}

type ContextPool struct {
	sync.Pool
}

func (ctx *HttpContext) reset() {
	ctx.Request.Body = ctx.Request.Body[0:0]
	ctx.Request.CallID = uint64(0)
	ctx.Response.Body = ctx.Response.Body[0:0]
	ctx.Response.Code = int32(0)
	ctx.Response.SessionID = uint64(0)
}

func NewContext() *HttpContext {
	return &HttpContext{
		Request: Request{
			Body:   make([]byte, 4096),
			CallID: 0,
		},
		Response: Response{
			Body:      make([]byte, 0),
			SessionID: 0,
			Code:      200,
		},
	}
}

func (p *ContextPool) GetContext() *HttpContext {
	service, _ := p.Get().(*HttpContext)
	return service
}

func (p *ContextPool) PutContext(service *HttpContext) {
	service.reset()
	p.Put(service)
}

func NewContextPool() ContextPool {
	return ContextPool{
		sync.Pool{
			New: func() interface{} { return NewContext() },
		}}
}

func (r *Request) Build(req *http.Request) error {
	if req == nil {
		return errors.New("natsproxy: Request cannot be nil")
	}

	buf := bytes.NewBuffer(r.Body)
	buf.Reset()
	if req.Body != nil {
		if _, err := io.Copy(buf, req.Body); err != nil {
			return err
		}
		if err := req.Body.Close(); err != nil {
			return err
		}
	}

	r.Body = buf.Bytes()
	return nil
}