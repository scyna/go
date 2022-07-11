package application

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	proto "github.com/scyna/go/manager/.proto/generated"
	"github.com/scyna/go/manager/model"
	"github.com/scyna/go/manager/repository"
	"github.com/scyna/go/scyna"
)

func CreateApplication(s *scyna.Service, request *proto.Application) {
	s.Logger.Info("Receive CreateApplicationRequest")

	if validateCreateApplication(request) != nil {
		s.Error(scyna.REQUEST_INVALID)
		return
	}

	var app model.Application
	app.FromDTO(request)

	if err := repository.CreateApplication(s.Logger, &app); err != nil {
		s.Error(err)
		return
	}

	s.Done(scyna.OK)
}

func validateCreateApplication(request *proto.Application) error {
	return validation.ValidateStruct(request,
		validation.Field(&request.Code, validation.Required, validation.Length(5, 100)),
		validation.Field(&request.Name, validation.Required, validation.Length(5, 200)),
		validation.Field(&request.OrgCode, validation.Required, validation.Length(5, 100)),
		validation.Field(&request.AuthPath, validation.Required, is.URL),
	)
}
