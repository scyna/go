package nats_manager

import (
	"fmt"

	proto "github.com/scyna/go/manager/.proto/generated"
	"github.com/scyna/go/scyna"
)

func DeleteStream(s *scyna.Service, request *proto.DeleteStreamRequest) {
	s.Logger.Info(fmt.Sprintf("%s\n", request.String()))

	if err := scyna.JetStream.DeleteStream(request.Name); err != nil {
		s.Error(scyna.SERVER_ERROR)
		s.Logger.Error(err.Error())
		return
	}

	s.Done(scyna.OK)
}
