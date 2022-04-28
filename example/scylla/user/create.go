package user

import (
	"github.com/scyna/go/example/scylla/proto"
	"github.com/scyna/go/scyna"
)

func Create(ctx *scyna.Context, request *proto.CreateUserRequest) {
	ctx.LOG.Info("Receive CreateRequest")
	/*TODO validate*/
	if err := Repository.Exist(ctx.LOG, request.User.Email); err == nil {
		ctx.Error(USER_EXISTED)
		return
	}

	user := FromProto(request.User)
	user.ID = scyna.ID.Next()
	if err := Repository.Create(ctx.LOG, user); err != nil {
		ctx.Error(err)
		return
	}
	ctx.Done(&proto.CreateUserResponse{Id: uint64(user.ID)})
}
