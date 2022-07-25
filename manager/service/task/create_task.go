package task

import (
	validation "github.com/go-ozzo/ozzo-validation"
	proto "github.com/scyna/go/manager/.proto/generated"
	"github.com/scyna/go/manager/model"
	"github.com/scyna/go/manager/repository"
	"github.com/scyna/go/manager/utils"
	"github.com/scyna/go/scyna"
)

func CreateTask(s *scyna.Service, request *proto.Task) {
	s.Logger.Info("Receive CreateEventRequest")

	if err := validation.ValidateStruct(request,
		validation.Field(&request.Module, validation.Required, validation.Length(1, 255), validation.Match(utils.NAME_PATTERN)),
		validation.Field(&request.Task, validation.Required, validation.Length(1, 255), validation.Match(utils.NAME_PATTERN)),
	); err != nil {
		s.Error(scyna.REQUEST_INVALID)
		return
	}

	if !repository.CheckModule(request.Module) {
		s.Error(model.MODULE_NOT_EXIST)
		return
	}

	if err := utils.CreateTaskConsumer(request.Module, request.Task); err != nil {
		s.Error(model.CAN_NOT_CREATE_CONSUMER)
		return
	}

	s.Done(scyna.OK)
}
