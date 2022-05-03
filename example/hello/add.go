package hello

import (
	"github.com/scyna/go/example/hello/proto"
	"github.com/scyna/go/scyna"
)

func Add(ctx *scyna.Context, request *proto.AddRequest) {
	ctx.LOG.Info("Receive AddRequest")

	sum := request.A + request.B
	if sum > 100 {
		ctx.Error(ADD_RESULT_TOO_BIG)
		return
	}

	ctx.Done(&proto.AddResponse{Sum: sum})
}
