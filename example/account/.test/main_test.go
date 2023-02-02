package test

import (
	"github.com/scyna/go/example/account/model"
	"github.com/scyna/go/example/account/service"
	"os"
	"testing"
	"time"

	"github.com/scyna/go/scyna"
	scyna_test "github.com/scyna/go/scyna/testing"
)

func TestMain(m *testing.M) {
	scyna_test.Init()

	scyna.RegisterService(model.CREATE_ACCOUNT_URL, service.CreateAccount)
	scyna.RegisterService(model.GET_ACCOUNT_BY_EMAIl_URL, service.GetAccountByEmail)

	exitVal := m.Run()
	time.Sleep(100 * time.Millisecond)
	scyna_test.Release()
	os.Exit(exitVal)
}
