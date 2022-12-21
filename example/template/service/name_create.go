package service

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/scyna/go/scyna"
	proto "vf.support/.proto/generated"
	"vf.support/model"
	"vf.support/repository"
)

func CreateName(s *scyna.Service, request *proto.CreateNameRequest) {
	if err := validation.ValidateStruct(
		request,
		validation.Field(&request.Name, validation.Required),
	); err != nil {
		s.Error(scyna.REQUEST_INVALID)
		s.Logger.Error(err.Error())
		return
	}

	id := scyna.ID.Next()
	if err := repository.CreateName(s.Logger, request.Name); err != nil {
		s.Error(err)
		return
	}

	s.Done(scyna.OK)

	s.PostEventAndActivity(model.NAME_CREATED_CHANNEL, &proto.NameCreated{
		Name: request.Name,
	}, []uint64{
		id,
	})
}
