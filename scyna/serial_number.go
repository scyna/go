package scyna

import (
	"fmt"
	"sync"
)

type serialNumber struct {
	key    string
	mutex  sync.Mutex
	prefix uint32
	last   uint64
	next   uint64
}

func InitSerialNumber(key string) *serialNumber {
	return &serialNumber{
		key:    key,
		prefix: 0,
		last:   0,
		next:   0,
	}
}

func (sn *serialNumber) Next() string {
	sn.mutex.Lock()
	defer sn.mutex.Unlock()

	if sn.next < sn.last {
		sn.next++
	} else {
		request := GetSNRequest{Key: sn.key}
		var response GetSNResponse
		if r := callService(GEN_GET_SN_URL, &request, &response); r.Code == 0 {
			sn.prefix = response.Prefix
			sn.next = response.Start
			sn.last = response.End
		} else {
			Fatal("Can not get SerialNumber")
		}
	}
	return fmt.Sprintf("%d%07d", sn.prefix, sn.next)
}
