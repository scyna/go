package organization

import (
	proto "github.com/scyna/go/manager/.proto/generated"
	"github.com/scyna/go/scyna"
)

func CreateModule(s *scyna.Service, request *proto.Module) {
	s.Logger.Info("Receive CreateModuleRequest")

	/*TODO*/

	s.Done(scyna.OK)
}
