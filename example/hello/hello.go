package hello

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/scyna/go/example/hello/proto"
	"github.com/scyna/go/scyna"
)

func Hello(s *scyna.Service, request *proto.HelloRequest) {
	s.Logger.Info("Receive HelloRequest")

	if err := validateHelloRequest(request); err != nil {
		s.Error(scyna.REQUEST_INVALID)
		return
	}

	s.Done(&proto.HelloResponse{Content: "Hello " + request.Name})
}

func validateHelloRequest(request *proto.HelloRequest) error {
	return validation.Validate(request.Name, validation.Required, validation.Length(3, 40))
}
