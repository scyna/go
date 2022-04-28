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
	scyna.RegisterStatelessService(basic.HELLO_URL, basic.Hello)
	scyna.RegisterStatefulService(basic.ECHO_URL, basic.Echo)
	scyna.RegisterStatefulService(basic.ADD_URL, basic.Add)

	/*register signals*/
	scyna.RegisterStatelessSignal(basic.STATELESS_CHANNEL, basic.StatelessSignal)
	scyna.RegisterStatefulSignal(basic.TEST_SIGNAL_CHANNEL, basic.TestSignal)

	exitVal := m.Run()
	scyna_test.Release()
	os.Exit(exitVal)
}
