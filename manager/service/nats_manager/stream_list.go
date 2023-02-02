package nats_manager

import (
	"context"
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
	proto "github.com/scyna/go/manager/.proto/generated"
	"github.com/scyna/go/scyna"
)

func ListStream(s *scyna.Service, request *proto.ListStreamRequest) {
	s.Logger.Info(fmt.Sprintf("%s\n", request.String()))

	ctx, cancel := context.WithTimeout(context.Background(), 9*time.Second)
	defer cancel()
	var response proto.ListStreamResponse
	var i = 0

	for info := range scyna.JetStream.Streams(nats.Context(ctx)) {
		i++
		response.Streams = append(response.Streams, &proto.StreamInfo{
			Name:    info.Config.Name,
			Subject: info.Config.Subjects,
			Created: info.Created.Format(time.RFC3339),
		})
		if i == 20 {
			break
		}
	}

	s.Done(&response)
}
