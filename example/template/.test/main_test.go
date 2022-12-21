package test

import (
	"os"
	"testing"
	"time"

	"github.com/scyna/go/scyna"
	scyna_test "github.com/scyna/go/scyna/testing"
	"vf.support/model"
	"vf.support/service"
)

func TestMain(m *testing.M) {
	scyna_test.Init()

	scyna.RegisterService(model.CREATE_NAME_URL, service.CreateName)

	exitVal := m.Run()
	time.Sleep(100 * time.Millisecond)
	scyna_test.Release()
	os.Exit(exitVal)
}
