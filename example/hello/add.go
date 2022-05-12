package hello

import (
	"github.com/scyna/go/example/hello/proto"
	"github.com/scyna/go/scyna"
)

func Add(s *scyna.Service, request *proto.AddRequest) {
	s.Logger.Info("Receive AddRequest")

	sum := request.A + request.B
	if sum > 100 {
		s.Error(ADD_RESULT_TOO_BIG)
		return
	}

	s.Done(&proto.AddResponse{Sum: sum})
}
