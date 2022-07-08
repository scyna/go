package organization

import (
	proto "github.com/scyna/go/manager/.proto/generated"
	"github.com/scyna/go/scyna"
)

func CreateModule(s *scyna.Service, request *proto.Module) {
	s.Logger.Info("Receive CreateModuleRequest")

	/*TODO: check input*/
	/*TODO: save module to database */
	/*TODO: create stream on NATS for sync */
	/*TODO: create stream on NATS for event */

	s.Done(scyna.OK)
}
