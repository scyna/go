package organization

import (
	proto "github.com/scyna/go/manager/.proto/generated"
	"github.com/scyna/go/scyna"
)

func DestroyOrganization(s *scyna.Service, request *proto.DestroyOrganizationRequest) {
	s.Logger.Info("Receive DestroyOrganizationRequest")

	/*TODO*/

	s.Done(scyna.OK)
}
