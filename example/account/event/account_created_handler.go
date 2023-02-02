package event

import (
	proto "github.com/scyna/go/example/account/.proto/generated"
	"github.com/scyna/go/scyna"
)

func AccountCreatedHandler(ctx *scyna.Context, event *proto.AccountCreated) {
	//Handle
}
