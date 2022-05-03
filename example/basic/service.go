package basic

import (
	"github.com/scyna/go/example/basic/proto"
	"github.com/scyna/go/scyna"
)

func Hello(ctx *scyna.Context) {
	ctx.LOG.Info("Receive HelloCommand")
	ctx.Done(&proto.HelloResponse{Text: "Hello World"})
}

func Echo(ctx *scyna.Context, request *proto.EchoRequest) {
	ctx.LOG.Info("Receive EchoRequest")
	ctx.Done(&proto.EchoResponse{Text: request.Text})
}

func Add(ctx *scyna.Context, request *proto.AddRequest) {
	ctx.LOG.Info("Receive AddRequest")
	sum := request.A + request.B
	if sum > 100 {
		ctx.Error(scyna.REQUEST_INVALID)
	} else {
		ctx.Done(&proto.AddResponse{Sum: sum})
	}
}
