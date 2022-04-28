package user

import (
	"github.com/scyna/go/example/scylla/proto"
	"github.com/scyna/go/scyna"
)

func Get(ctx *scyna.Context, request *proto.GetUserRequest) {
	ctx.LOG.Info("Receive CreateRequest")
	if err, user := Repository.GetByEmail(ctx.LOG, request.Email); err != nil {
		ctx.Error(err)
		return
	} else {
		ctx.Done(&proto.GetUserResponse{User: user.ToProto()})
	}
}
