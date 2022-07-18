package module

import (
	"log"

	validation "github.com/go-ozzo/ozzo-validation"
	proto "github.com/scyna/go/manager/.proto/generated"
	"github.com/scyna/go/manager/model"
	"github.com/scyna/go/manager/repository"
	"github.com/scyna/go/manager/utils"
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

	if !utils.ValidatePassword(request.Secret) {
		s.Error(scyna.REQUEST_INVALID)
		return
	}

	var module model.Module
	module.FromDTO(request)
	hash, _ := utils.HashPassword(module.Secret)
	module.Secret = hash

	if err := repository.CreateModule(s.Logger, &module); err != nil {
		s.Error(err)
		return
	}

	if err := utils.CreateStreamForModule(module.Code); err != nil {
		log.Println(err.Error())
		s.Error(model.CAN_NOT_CREATE_STREAM)
		return
	}

	s.Done(scyna.OK)
}

func validateModule(request *proto.Module) error {
	return validation.ValidateStruct(request,
		validation.Field(&request.Organization, validation.Required, validation.Length(1, 100)),
		validation.Field(&request.Code, validation.Required, validation.Length(3, 200), validation.Match(utils.NAME_PATTERN)),
		validation.Field(&request.Secret, validation.Required, validation.Length(5, 20)),
	)
}
