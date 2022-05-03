package basic

import (
	"github.com/scyna/go/example/basic/proto"
	"github.com/scyna/go/scyna"
)

func HelloSignal(LOG scyna.Logger, signal *proto.HelloSignal) {
	LOG.Info("Received TestSignal:" + signal.Text)
}
