package user_test

import (
	"testing"

	"github.com/scyna/go/example/scylla/proto"
	"github.com/scyna/go/example/scylla/user"
	scyna_test "github.com/scyna/go/scyna/testing"
)

func TestCreateShouldReturnSuccess(t *testing.T) {
	cleanup()
	requestCreate := &proto.CreateUserRequest{
		User: &proto.User{
			Email:    "a@gmail.com",
			Name:     "Nguyen Van A",
			Password: "1234565",
		},
	}

	var responseCreate proto.CreateUserRequest
	scyna_test.CallServiceParseResponse(t, user.CREATE_USER_URL, requestCreate, &responseCreate, 200)
}
