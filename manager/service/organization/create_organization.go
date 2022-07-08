package organization

import (
	proto "github.com/scyna/go/manager/.proto/generated"
	"github.com/scyna/go/scyna"
)

func CreateOrganization(s *scyna.Service, request *proto.Organization) {
	s.Logger.Info("Receive CreateOrganizationRequest")

	/*TODO*/

	s.Done(scyna.OK)
}
