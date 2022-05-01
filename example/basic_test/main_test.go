package basic_test

import (
	"os"
	"testing"

	"github.com/scyna/go/example/basic"
	"github.com/scyna/go/scyna"
	scyna_test "github.com/scyna/go/scyna/testing"
)

func TestMain(m *testing.M) {
	scyna_test.Init()

	/*register services*/
	scyna.RegisterServiceLite(basic.HELLO_URL, basic.Hello)
	scyna.RegisterService(basic.ECHO_URL, basic.Echo)
	scyna.RegisterService(basic.ADD_URL, basic.Add)

	/*register signals*/
	scyna.RegisterSignalLite(basic.STATELESS_CHANNEL, basic.StatelessSignal)
	scyna.RegisterSignal(basic.HELLO_SIGNAL_CHANNEL, basic.HelloSignal)

	exitVal := m.Run()
	scyna_test.Release()
	os.Exit(exitVal)
}
