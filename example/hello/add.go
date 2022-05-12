package hello

import (
	"github.com/scyna/go/example/hello/proto"
	"github.com/scyna/go/scyna"
)

func Add(c *scyna.Service, request *proto.AddRequest) {
	c.Logger.Info("Receive AddRequest")

	sum := request.A + request.B
	if sum > 100 {
		c.Error(ADD_RESULT_TOO_BIG)
		return
	}

	c.Done(&proto.AddResponse{Sum: sum})
}
