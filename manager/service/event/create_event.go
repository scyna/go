package organization

import (
	proto "github.com/scyna/go/manager/.proto/generated"
	"github.com/scyna/go/scyna"
)

func CreateEvent(s *scyna.Service, request *proto.Event) {
	s.Logger.Info("Receive CreateEventRequest")

	/*TODO*/

	s.Done(scyna.OK)
}
