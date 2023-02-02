package nats_manager

import (
	"fmt"

	proto "github.com/scyna/go/manager/.proto/generated"
	"github.com/scyna/go/scyna"
)

func GetStream(s *scyna.Service, request *proto.DeleteStreamRequest) {
	s.Logger.Info(fmt.Sprintf("%s\n", request.String()))

	// TODO
	s.Done(scyna.OK)
}
