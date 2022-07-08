package organization

import (
	proto "github.com/scyna/go/manager/.proto/generated"
	"github.com/scyna/go/scyna"
)

func CreateSync(s *scyna.Service, request *proto.Sync) {
	s.Logger.Info("Receive CreateSyncRequest")

	/*TODO*/

	s.Done(scyna.OK)
}
