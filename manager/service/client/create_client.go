package organization

import (
	proto "github.com/scyna/go/manager/.proto/generated"
	"github.com/scyna/go/scyna"
)

func CreateClient(s *scyna.Service, request *proto.Client) {
	s.Logger.Info("Receive CreateClientRequest")

	/*TODO*/

	s.Done(scyna.OK)
}
