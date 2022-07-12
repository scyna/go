package sync

import (
	proto "github.com/scyna/go/manager/.proto/generated"
	"github.com/scyna/go/manager/model"
	"github.com/scyna/go/manager/repository"
	"github.com/scyna/go/manager/utils"
	"github.com/scyna/go/scyna"
)

func CreateSync(s *scyna.Service, request *proto.CreateSyncRequest) {
	s.Logger.Info("Receive CreateSyncRequest")

	if !repository.CheckModule(request.Module) {
		s.Error(model.MODULE_NOT_EXIST)
		return
	}

	/*TODO: validate channel*/

	consumer := scyna.GetSyncConsumer(request.Module, request.Channel, request.Receiver)

	if err := utils.AddConsumer(request.Module, consumer, "SYNC", request.Channel, request.Module+"."+request.Channel); err != nil {
		s.Error(model.CAN_NOT_CREATE_CONSUMER)
		return
	}
	s.Done(scyna.OK)
}
