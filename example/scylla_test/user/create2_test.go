package user_test

import (
	"testing"

	"github.com/scyna/go/example/scylla/proto"
	"github.com/scyna/go/example/scylla/user"
	"github.com/scyna/go/scyna"
	scyna_test "github.com/scyna/go/scyna/testing"
)

func TestCreate2ShouldReturnSuccess(t *testing.T) {
	cleanup()
	scyna_test.ServiceTest(user.CREATE_USER_URL).
		WithRequest(&proto.CreateUserRequest{User: &proto.User{
			Email:    "a@gmail.com",
			Name:     "Nguyen Van A",
			Password: "1234565",
		}}).
		ExpectOK().Run(t)

	scyna_test.ServiceTest(user.GET_USER_URL).
		WithRequest(&proto.GetUserRequest{Email: "a@gmail.com"}).
		ExpectOK().Run(t)
}

func TestCreate2BadEmail(t *testing.T) {
	cleanup()
	scyna_test.ServiceTest(user.CREATE_USER_URL).
		WithRequest(&proto.CreateUserRequest{User: &proto.User{
			Email:    "a+gmail.com",
			Name:     "Nguyen Van A",
			Password: "1234565",
		}}).
		ExpectError(scyna.REQUEST_INVALID).Run(t)
}
