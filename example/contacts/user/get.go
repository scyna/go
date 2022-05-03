package user

import (
	"github.com/scyna/go/example/contacts/proto"
	"github.com/scyna/go/scyna"
)

func Get(ctx *scyna.Context, request *proto.GetUserRequest) {
	ctx.LOG.Info("Receive GetUserRequest")
	if err, user := Repository.GetByEmail(ctx.LOG, request.Email); err != nil {
		ctx.Error(err)
		return
	} else {
		ctx.Done(&proto.GetUserResponse{User: user.ToProto()})
	}
}
