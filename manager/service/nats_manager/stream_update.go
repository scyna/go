package nats_manager

import (
	"fmt"

	proto "github.com/scyna/go/manager/.proto/generated"
	"github.com/scyna/go/scyna"
)

func UpdateStream(s *scyna.Service, request *proto.UpdateStreamRequest) {
	s.Logger.Info(fmt.Sprintf("%s\n", request.String()))
	// TODO
	s.Done(scyna.OK)
}
