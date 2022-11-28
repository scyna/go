package user

import (
	"github.com/scyna/go/example/contacts/proto"
	"github.com/scyna/go/scyna"
)

const CreateUserUrl = "/scyna.example/user/create"

func CreateUserHandler(cmd *scyna.Command, request *proto.User) {
	cmd.Logger.Info("Receive CreateUserRequest")
	if err := validateCreateUserRequest(request); err != nil {
		cmd.Error(scyna.REQUEST_INVALID)
		return
	}

	if err, _ := Repository.GetByEmail(cmd.Logger, request.Email); err == nil {
		cmd.Error(USER_EXISTED)
		return

	}

	user := FromDTO(request)
	user.ID = scyna.ID.Next()

	Repository.PrepareCreate(cmd, user)

	cmd.Done(&proto.CreateUserResponse{Id: user.ID},
		user.ID,
		"ex.user.user_created",
		&proto.UserCreated{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email})
}
