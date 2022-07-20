package user

import (
	"github.com/scyna/go/example/contacts/proto"
	"github.com/scyna/go/scyna"
	"log"
	"net/http"
)

func HandlerEventMessage(ctx *scyna.Context, message *proto.User) {
	log.Println("X Event")
}

func HandlerSyncMessage(ctx *scyna.Context, message *proto.User) *http.Request {
	log.Println("X Sync " + message.Email)
	return nil
}
