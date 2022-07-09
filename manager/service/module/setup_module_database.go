package module

import (
	proto "github.com/scyna/go/manager/.proto/generated"
	"github.com/scyna/go/manager/model"
	"github.com/scyna/go/manager/repository"
	"github.com/scyna/go/scyna"
)

func SetupModuleDatabase(s *scyna.Service, request *proto.SetupModuleDatabaseRequest) {
	s.Logger.Info("Receive SetupDatabaseRequest")

	if !repository.CheckModule(request.Module) {
		s.Error(model.MODULE_NOT_EXIST)
		return
	}

	/*TODO: create module keyspace */
	/*TODO: create EventStore */

	s.Done(scyna.OK)
}
