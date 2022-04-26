package scyna

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"sync"
)

type context struct {
	Request  Request
	Response Response
}

type ContextPool struct {
	sync.Pool
}

func (service *context) reset() {
	service.Request.Body = service.Request.Body[0:0]
	service.Request.CallID = uint64(0)
	service.Response.Body = service.Response.Body[0:0]
	service.Response.Code = int32(0)
	service.Response.SessionID = uint64(0)
}

func NewService() *context {
	return &context{
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

func (p *ContextPool) GetService() *context {
	service, _ := p.Get().(*context)
	return service
}

func (p *ContextPool) PutService(service *context) {
	service.reset()
	p.Put(service)
}

func NewServicePool() ContextPool {
	return ContextPool{
		sync.Pool{
			New: func() interface{} { return NewService() },
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
