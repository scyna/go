package basic

import "github.com/scyna/go/scyna"

func Hello(ctx *scyna.Context) {
	ctx.LOG.Info("Receive HelloRequest")
	ctx.Done(&HelloResponse{Text: "Hello World"})
}

func Echo(ctx *scyna.Context, request *EchoRequest) {
	ctx.LOG.Info("Receive EchoRequest")
	ctx.Done(&EchoResponse{Text: request.Text})
}

func Add(ctx *scyna.Context, request *AddRequest) {
	ctx.LOG.Info("Receive AddRequest")
	sum := request.A + request.B
	if sum > 100 {
		ctx.Error(scyna.REQUEST_INVALID)
	} else {
		ctx.Done(&AddResponse{Sum: sum})
	}
}
