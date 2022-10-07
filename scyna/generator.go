package scyna

import (
	"sync"
	"time"

	"google.golang.org/protobuf/proto"
)

type generator struct {
	mutex  sync.Mutex
	prefix uint32
	last   uint64
	next   uint64
}

func (g *generator) Reset(prefix uint32, last uint64, next uint64) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	g.prefix = prefix
	g.last = last
	g.next = next
}

func (g *generator) Next() uint64 {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	if g.next < g.last {
		g.next++
	} else {
		if !g.getID() {
			Fatal("Can not create generator")
		}
	}
	return (uint64(g.prefix) << 44) + g.next
}

func (g *generator) getID() bool {
	req := Request{TraceID: 0, JSON: false}
	res := Response{}

	if data, err := proto.Marshal(&req); err == nil {
		if msg, err := Connection.Request(PublishURL(GEN_GET_ID_URL), data, REQUEST_TIMEOUT*time.Second); err == nil {
			if err := proto.Unmarshal(msg.Data, &res); err == nil {
				if res.Code == 200 {
					var response GetIDResponse
					if err := proto.Unmarshal(res.Body, &response); err == nil {
						g.prefix = response.Prefix
						g.next = response.Start
						g.last = response.End
						return true
					}
				}
			}
		}
	}
	return false
}
