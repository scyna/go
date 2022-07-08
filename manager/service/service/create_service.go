package organization

import (
	proto "github.com/scyna/go/manager/.proto/generated"
	"github.com/scyna/go/scyna"
)

func CreateService(s *scyna.Service, request *proto.Service) {
	s.Logger.Info("Receive CreateServiceRequest")

	/*TODO*/

	s.Done(scyna.OK)
}
