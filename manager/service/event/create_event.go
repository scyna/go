package event

import (
	validation "github.com/go-ozzo/ozzo-validation"
	proto "github.com/scyna/go/manager/.proto/generated"
	"github.com/scyna/go/manager/model"
	"github.com/scyna/go/manager/repository"
	"github.com/scyna/go/manager/utils"
	"github.com/scyna/go/scyna"
)

func CreateEvent(s *scyna.Service, request *proto.CreateEventRequest) {
	s.Logger.Info("Receive CreateEventRequest")

	if validation.Validate(request.Channel, validation.Match(utils.NAME_PATTERN)) != nil {
		s.Error(scyna.REQUEST_INVALID)
		return
	}

	if !repository.CheckModule(request.SenderModule) {
		s.Error(model.MODULE_NOT_EXIST)
		return
	}

	if !repository.CheckModule(request.ReceiverModule) {
		s.Error(model.MODULE_NOT_EXIST)
		return
	}

	if err := utils.CreateConsumer(request.SenderModule, request.ReceiverModule); err != nil {
		s.Error(model.CAN_NOT_CREATE_CONSUMER)
		return
	}

	s.Done(scyna.OK)
}
