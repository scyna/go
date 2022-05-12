package user

import (
	"github.com/scyna/go/example/contacts/proto"
	"github.com/scyna/go/scyna"
)

func Get(s *scyna.Service, request *proto.GetUserRequest) {
	s.Logger.Info("Receive GetUserRequest")
	if err, user := Repository.GetByEmail(s.Logger, request.Email); err != nil {
		s.Error(err)
		return
	} else {
		s.Done(&proto.GetUserResponse{User: user.ToDTO()})
	}
}
