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

	/*register APIs*/
	scyna.RegisterStatelessService("/example/basic/hello", basic.Hello)
	scyna.RegisterStatefulService("/example/basic/echo", basic.Echo)
	scyna.RegisterStatefulService("/example/basic/add", basic.Add)

	exitVal := m.Run()
	scyna_test.Release()
	os.Exit(exitVal)
}
