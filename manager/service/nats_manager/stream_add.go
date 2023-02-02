package nats_manager

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
	proto "github.com/scyna/go/manager/.proto/generated"
	"github.com/scyna/go/scyna"
)

func AddStream(s *scyna.Service, request *proto.AddStreamRequest) {
	s.Logger.Info(fmt.Sprintf("%s\n", request.String()))

	if _, err := scyna.JetStream.AddStream(&nats.StreamConfig{
		Name:         request.Name,
		Subjects:     []string{request.Name + ".>"},
		Storage:      nats.FileStorage,
		MaxAge:       time.Hour * 24 * 7, //keep for a week
		Replicas:     3,
		Retention:    nats.LimitsPolicy,
		MaxMsgs:      -1,
		MaxConsumers: -1,
		MaxBytes:     -1,
		MaxMsgSize:   -1,
	}); err != nil {
		s.Error(scyna.SERVER_ERROR)
		s.Logger.Error(err.Error())
		return
	}

	s.Done(scyna.OK)
}
