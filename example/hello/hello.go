package hello

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/scyna/go/example/hello/proto"
	"github.com/scyna/go/scyna"
)

func Hello(c *scyna.Service, request *proto.HelloRequest) {
	c.Logger.Info("Receive HelloRequest")

	if err := validateHelloRequest(request); err != nil {
		c.Error(scyna.REQUEST_INVALID)
		return
	}

	c.Done(&proto.HelloResponse{Content: "Hello " + request.Name})
}

func validateHelloRequest(request *proto.HelloRequest) error {
	return validation.Validate(request.Name, validation.Required, validation.Length(3, 40))
}
