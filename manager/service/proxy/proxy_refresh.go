package proxy

import (
	proto "github.com/scyna/go/manager/.proto/generated"
	"github.com/scyna/go/scyna"
)

func Refresh(s *scyna.Service, request *proto.RefreshProxy) {
	s.Logger.Info("Refresh proxy")
	s.Done(scyna.OK)
	scyna.Connection.Publish(scyna.CLIENT_UPDATE_CHANNEL, []byte("Reload proxy"))
}
