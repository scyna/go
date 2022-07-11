package client

import (
	validation "github.com/go-ozzo/ozzo-validation"
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

	if !repository.CheckClient(request.Organization, request.Id) {
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
	return validation.ValidateStruct(request,
		validation.Field(&request.Id, validation.Required, validation.Length(5, 100)),
		validation.Field(&request.Service, validation.Required, validation.Length(8, 50)),
		validation.Field(&request.Organization, validation.Required, validation.Length(8, 50)),
	)
}
