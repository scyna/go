package organization

import (
	proto "github.com/scyna/go/manager/.proto/generated"
	"github.com/scyna/go/scyna"
)

func CreateApplication(s *scyna.Service, request *proto.Application) {
	s.Logger.Info("Receive CreateApplicationRequest")

	/*TODO*/

	s.Done(scyna.OK)
}
