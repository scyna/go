package service

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	proto "github.com/scyna/go/manager/.proto/generated"
	"github.com/scyna/go/manager/model"
	"github.com/scyna/go/manager/repository"
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
		validation.Field(&request.Url, validation.Required, is.URL),
	)
}
