package organization

import (
	validation "github.com/go-ozzo/ozzo-validation"
	proto "github.com/scyna/go/manager/.proto/generated"
	"github.com/scyna/go/manager/model"
	"github.com/scyna/go/manager/repository"
	"github.com/scyna/go/scyna"
)

func CreateOrganization(s *scyna.Service, request *proto.Organization) {
	s.Logger.Info("Receive CreateOrganizationRequest")

	if validateOrganization(request) != nil {
		s.Error(scyna.REQUEST_INVALID)
		return
	}

	var org model.Organization
	org.FromDTO(request)

	if err := repository.CreateOrganization(&org); err != nil {
		s.Error(err)
		return
	}

	s.Done(scyna.OK)
}

func validateOrganization(request *proto.Organization) error {
	return validation.ValidateStruct(request,
		validation.Field(&request.Code, validation.Required, validation.Length(1, 100)), //FIXME: name rules
		validation.Field(&request.Name, validation.Required, validation.Length(1, 200)),
		validation.Field(&request.Password, validation.Required),
	)
}
