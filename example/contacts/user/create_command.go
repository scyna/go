package user

import (
	"github.com/scyna/go/example/contacts/proto"
	"github.com/scyna/go/scyna"
)

func CreateUserHandler(s *scyna.Command, request *proto.User) {
	s.Logger.Info("Receive CreateUserRequest")
	if err := validateCreateUserRequest(request); err != nil {
		s.Error(scyna.REQUEST_INVALID)
		return
	}

	if err, _ := Repository.GetByEmail(s.Logger, request.Email); err == nil {
		s.Error(USER_EXISTED)
		return
	}

	user := FromDTO(request)
	user.ID = scyna.ID.Next()

	//if err := Repository.Create(s.Logger, user); err != nil {
	//	s.Error(err)
	//	return
	//}

	//s.PostSync("account", user.ToDTO())

	s.Done(&proto.CreateUserResponse{Id: user.ID}, user.ID, "ex.user.user_created", nil) //FIXME
}
