package module

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
	proto "github.com/scyna/go/manager/.proto/generated"
	"github.com/scyna/go/manager/model"
	"github.com/scyna/go/manager/repository"
	"github.com/scyna/go/manager/utils"
	"github.com/scyna/go/scyna"
	"log"
	"regexp"
	"strings"
)

func CreateModule(s *scyna.Service, request *proto.Module) {
	s.Logger.Info("Receive CreateModuleRequest")

	if validateModule(request) != nil {
		s.Error(scyna.REQUEST_INVALID)
		return
	}

	if err := validateModuleCode(request.Code); err != nil {
		s.Error(model.MODULE_CODE_BAD_FORMAT)
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
	streamName := strings.Replace(module.Code, ".", "_", -1)
	if err := utils.CreateStream(streamName); err != nil {
		log.Println(err.Error())
		s.Error(model.CAN_NOT_CREATE_STREAM)
		return
	}

	s.Done(scyna.OK)
}

func validateModule(request *proto.Module) error {
	return validation.ValidateStruct(request,
		validation.Field(&request.Organization, validation.Required, validation.Length(1, 100)),
		validation.Field(&request.Code, validation.Required, validation.Length(5, 100)), //FIXME: module name rules
		validation.Field(&request.Secret, validation.Required, validation.Length(5, 20)),
	)
}

func validateModuleCode(value interface{}) error {
	dob, _ := value.(string)
	regex := "^[a-z0-9_]*$" // module_name
	match, _ := regexp.MatchString(regex, dob)
	if !match {
		return errors.New("invalid module code")
	}
	return nil
}
