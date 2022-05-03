package hello_test

import (
	"os"
	"testing"

	"github.com/scyna/go/example/hello"
	"github.com/scyna/go/scyna"
	scyna_test "github.com/scyna/go/scyna/testing"
)

func TestMain(m *testing.M) {
	scyna_test.Init()

	scyna.RegisterService(hello.HELLO_URL, hello.Hello)
	scyna.RegisterService(hello.ADD_URL, hello.Add)

	exitVal := m.Run()
	scyna_test.Release()
	os.Exit(exitVal)
}
