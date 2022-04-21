package scyna

import (
	"log"
	sync "sync"
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
			log.Fatal("Can not create generator")
		}
	}
	return (uint64(g.prefix) << 44) + g.next
}

func (g *generator) getID() bool {
	callID := uint64(0)

	var req Request
	var res Response

	req.CallID = callID
	req.JSON = false

	reqBytes, err := proto.Marshal(&req)
	if err != nil {
		return false
	}

	msg, respErr := Connection.Request(PublishURL(GEN_GET_ID_URL), reqBytes, 10*time.Second)
	if respErr != nil {
		return false
	}

	err = res.ReadFrom(msg.Data)
	if err != nil {
		return false
	}

	if res.Code == 200 {
		var response GetIDResponse
		if err := proto.Unmarshal(res.Body, &response); err == nil {
			g.prefix = response.Prefix
			g.next = response.Start
			g.last = response.End
			return true
		}
	}
	return false
}
