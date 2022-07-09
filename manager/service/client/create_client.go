package client

import (
	validation "github.com/go-ozzo/ozzo-validation"
	proto "github.com/scyna/go/manager/.proto/generated"
	"github.com/scyna/go/manager/model"
	"github.com/scyna/go/manager/repository"
	"github.com/scyna/go/scyna"
)

func CreateClient(s *scyna.Service, request *proto.Client) {
	s.Logger.Info("Receive CreateClientRequest")

	if validateClient(request) != nil {
		s.Error(scyna.REQUEST_INVALID)
		return
	}

	if !repository.CheckOrganization(request.Organization) {
		s.Error(model.ORGANIZATION_NOT_EXIST)
		return
	}

	var client model.Client
	client.FromDTO(request)
	if err := repository.CreateClient(&client); err != nil {
		s.Error(err)
		return
	}

	s.Done(scyna.OK)
}

func validateClient(request *proto.Client) error {
	return validation.ValidateStruct(request,
		validation.Field(&request.Id, validation.Required, validation.Length(1, 100)),     //FIXME: name rules
		validation.Field(&request.Secret, validation.Required, validation.Length(1, 200)), //FIXME: secret rules
	)
}
