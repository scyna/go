package user_test

import (
	"os"
	"testing"

	"github.com/scyna/go/example/contacts/user"
	"github.com/scyna/go/scyna"
	scyna_test "github.com/scyna/go/scyna/testing"
)

func TestMain(m *testing.M) {
	scyna_test.Init()

	/*register services*/
	scyna.RegisterService(user.CREATE_USER_URL, user.CreateUser)
	scyna.RegisterService(user.GET_USER_URL, user.GetUserByEmail)

	exitVal := m.Run()
	cleanup()
	scyna_test.Release()
	os.Exit(exitVal)
}

func cleanup() {
	scyna.DB.Query("TRUNCATE ex.user", nil).ExecRelease()
}
