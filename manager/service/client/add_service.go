package client

import (
	proto "github.com/scyna/go/manager/.proto/generated"
	"github.com/scyna/go/manager/model"
	"github.com/scyna/go/manager/repository"
	"github.com/scyna/go/scyna"
)

func AddService(s *scyna.Service, request *proto.ClientAddServiceRequest) {
	s.Logger.Info("Receive ClientAddServiceRequest")

	if validateAddService(request) != nil {
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

	if err := repository.AddService(s.Logger, request.Id, request.Service); err != nil {
		s.Error(err)
		return
	}

	s.Done(scyna.OK)

}

func validateAddService(request *proto.ClientAddServiceRequest) error {
	/*TODO*/
	return nil
}
