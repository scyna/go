package hello

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/scyna/go/example/hello/proto"
	"github.com/scyna/go/scyna"
)

func Hello(ctx *scyna.Context, request *proto.HelloRequest) {
	ctx.LOG.Info("Receive HelloRequest")

	if err := validateHelloRequest(request); err != nil {
		ctx.Error(scyna.REQUEST_INVALID)
		return
	}

	ctx.Done(&proto.HelloResponse{Content: "Hello " + request.Name})
}

func validateHelloRequest(request *proto.HelloRequest) error {
	return validation.ValidateStruct(request,
		validation.Field(&request.Name, validation.Required, validation.Length(3, 40)))
}
