package basic_test

import (
	"testing"
	"time"

	"github.com/scyna/go/example/basic"
	"github.com/scyna/go/example/basic/proto"
	"github.com/scyna/go/scyna"
)

func TestHelloSignal(t *testing.T) {
	scyna.EmitSignalLite(basic.STATELESS_CHANNEL)
	scyna.EmitSignal(basic.HELLO_SIGNAL_CHANNEL, &proto.HelloSignal{Text: "Hello"})
	time.Sleep(time.Second * 2)
}
