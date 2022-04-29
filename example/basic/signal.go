package basic

import (
	"github.com/scyna/go/example/basic/proto"
	"github.com/scyna/go/scyna"
)

func StatelessSignal(LOG scyna.Logger) {
	LOG.Info("Received StatelessSignal")
}

func HelloSignal(LOG scyna.Logger, signal *proto.HelloSignal) {
	LOG.Info("Received TestSignal:" + signal.Text)
}
