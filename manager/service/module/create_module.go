package module

import (
	validation "github.com/go-ozzo/ozzo-validation"
	proto "github.com/scyna/go/manager/.proto/generated"
	"github.com/scyna/go/manager/model"
	"github.com/scyna/go/manager/repository"
	"github.com/scyna/go/scyna"
)

func CreateModule(s *scyna.Service, request *proto.Module) {
	s.Logger.Info("Receive CreateModuleRequest")

	if validateModule(request) != nil {
		s.Error(scyna.REQUEST_INVALID)
		return
	}

	if !repository.CheckOrganization(request.Organization) {
		s.Error(model.ORGANIZATION_NOT_EXIST)
		return
	}

	if repository.CheckModule(request.Code) {
		s.Error(model.MODULE_EXISTED)
		return
	}

	var module model.Module
	module.FromDTO(request)

	if err := repository.CreateModule(&module); err != nil {
		s.Error(err)
		return
	}

	/*TODO: create stream on NATS for sync: `module_name.sync.*` */
	/*TODO: create stream on NATS for event `module_name.event.*` */

	s.Done(scyna.OK)
}

func validateModule(request *proto.Module) error {
	return validation.ValidateStruct(request,
		validation.Field(&request.Organization, validation.Required, validation.Length(1, 100)),
		validation.Field(&request.Code, validation.Required, validation.Length(1, 100)),
		validation.Field(&request.Secret, validation.Required, validation.Length(5, 20)),
	)
}
