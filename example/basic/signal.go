package basic

import (
	"log"

	"github.com/scyna/go/example/basic/proto"
	"github.com/scyna/go/scyna"
)

func StatelessSignal(LOG scyna.Logger) {
	log.Print("aaaaa")
	LOG.Info("Received StatelessSignal")
}

func TestSignal(LOG scyna.Logger, signal *proto.TestSignal) {
	LOG.Info("Received TestSignal:" + signal.Text)
}
