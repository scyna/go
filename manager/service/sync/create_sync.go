package sync

import (
	validation "github.com/go-ozzo/ozzo-validation"
	proto "github.com/scyna/go/manager/.proto/generated"
	"github.com/scyna/go/manager/model"
	"github.com/scyna/go/manager/repository"
	"github.com/scyna/go/manager/utils"
	"github.com/scyna/go/scyna"
)

func CreateSync(s *scyna.Service, request *proto.CreateSyncRequest) {
	s.Logger.Info("Receive CreateSyncRequest")

	if validation.Validate(request.Channel, validation.Match(utils.NAME_PATTERN)) != nil {
		s.Error(scyna.REQUEST_INVALID)
		return
	}

	if !repository.CheckModule(request.Module) {
		s.Error(model.MODULE_NOT_EXIST)
		return
	}

	if err := utils.CreateSyncConsumer(request.Module, request.Channel, request.Receiver); err != nil {
		s.Logger.Error(err.Error())
		s.Error(model.CAN_NOT_CREATE_CONSUMER)
		return
	}
	
	// TODO: save db

	s.Done(scyna.OK)
}
