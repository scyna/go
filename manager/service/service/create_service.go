package service

import (
	validation "github.com/go-ozzo/ozzo-validation"
	proto "github.com/scyna/go/manager/.proto/generated"
	"github.com/scyna/go/manager/model"
	"github.com/scyna/go/manager/repository"
	"github.com/scyna/go/manager/utils"
	"github.com/scyna/go/scyna"
)

func CreateService(s *scyna.Service, request *proto.Service) {
	s.Logger.Info("Receive CreateServiceRequest")

	if validateService(request) != nil {
		s.Error(scyna.REQUEST_INVALID)
		return
	}

	if !repository.CheckModule(request.Module) {
		s.Error(model.MODULE_NOT_EXIST)
		return
	}

	if err := utils.CheckPathService(request.Module, request.Url); err != nil {
		s.Logger.Warning(err.Error())
		s.Error(model.SERVICE_PATH_BAD_FORMAT)
		return
	}

	var service model.Service
	service.FromDTO(request)
	if err := repository.CreateService(s.Logger, &service); err != nil {
		s.Error(err)
		return
	}

	s.Done(scyna.OK)
}

func validateService(request *proto.Service) error {
	return validation.ValidateStruct(request,
		validation.Field(&request.Url, validation.Required, validation.Length(10, 100)),
	)
}
