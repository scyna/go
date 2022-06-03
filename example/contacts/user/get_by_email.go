package user

import (
	"github.com/scyna/go/example/contacts/proto"
	"github.com/scyna/go/scyna"
)

func GetUserByEmail(s *scyna.Service, request *proto.GetUserByEmailRequest) {
	s.Logger.Info("Receive GetUserRequest")
	if err, user := Repository.GetByEmail(s.Logger, request.Email); err != nil {
		s.Error(err)
		return
	} else {
		s.Done(user.ToDTO())
	}
}
