package module

import (
	proto "github.com/scyna/go/manager/.proto/generated"
	"github.com/scyna/go/scyna"
)

func ListModule(s *scyna.Service, request *proto.Module) {
	s.Logger.Info("Receive ListModuleRequest")
}
