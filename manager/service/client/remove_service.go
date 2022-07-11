package client

import (
	proto "github.com/scyna/go/manager/.proto/generated"
	"github.com/scyna/go/manager/model"
	"github.com/scyna/go/manager/repository"
	"github.com/scyna/go/scyna"
)

func RemoveService(s *scyna.Service, request *proto.ClientRemoveServiceRequest) {
	s.Logger.Info("Receive ClientRemoveServiceRequest")

	if validateRemoveService(request) != nil {
		s.Error(scyna.REQUEST_INVALID)
		return
	}

	if !repository.CheckOrganization(request.Organization) {
		s.Error(model.ORGANIZATION_NOT_EXIST)
		return
	}

	if !repository.CheckClient(request.Id) {
		s.Error(model.CLIENT_NOT_EXISTED)
		return
	}

	if err := repository.RemoveService(s.Logger, request.Id, request.Service); err != nil {
		s.Error(err)
		return
	}

	s.Done(scyna.OK)
}

func validateRemoveService(request *proto.ClientRemoveServiceRequest) error {
	/*TODO*/
	return nil
}
