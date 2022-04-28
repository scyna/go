package basic

import (
	"github.com/scyna/go/example/basic/proto"
	"github.com/scyna/go/scyna"
)

func StatelessSignal(LOG scyna.Logger) {
	LOG.Info("Received StatelessSignal")
}

func TestSignal(LOG scyna.Logger, signal *proto.TestSignal) {
	LOG.Info("Received TestSignal:" + signal.Text)
}
