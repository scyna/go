package organization

import (
	proto "github.com/scyna/go/manager/.proto/generated"
	"github.com/scyna/go/manager/model"
	"github.com/scyna/go/manager/repository"
	"github.com/scyna/go/manager/utils"
	"github.com/scyna/go/scyna"
)

func CreateEvent(s *scyna.Service, request *proto.CreateEventRequest) {
	s.Logger.Info("Receive CreateEventRequest")

	if !repository.CheckModule(request.SenderModule) {
		s.Error(model.MODULE_NOT_EXIST)
		return
	}

	if !repository.CheckModule(request.ReceiverModule) {
		s.Error(model.MODULE_NOT_EXIST)
		return
	}

	/*TODO: validate channel format*/

	consumer := scyna.GetEventConsumer(request.SenderModule, request.Channel, request.ReceiverModule)
	if err := utils.CreateConsumer(request.SenderModule, consumer, request.Channel); err != nil {
		s.Error(model.CAN_NOT_CREATE_CONSUMER)
		return
	}

	s.Done(scyna.OK)
}
