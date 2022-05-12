package user

import (
	"github.com/scyna/go/example/contacts/proto"
	"github.com/scyna/go/scyna"
)

func Get(ctx *scyna.Service, request *proto.GetUserRequest) {
	ctx.Logger.Info("Receive GetUserRequest")
	if err, user := Repository.GetByEmail(ctx.Logger, request.Email); err != nil {
		ctx.Error(err)
		return
	} else {
		ctx.Done(&proto.GetUserResponse{User: user.ToDTO()})
	}
}
