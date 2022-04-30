package user_test

import (
	"testing"

	"github.com/scyna/go/example/scylla/proto"
	"github.com/scyna/go/example/scylla/user"
	scyna_test "github.com/scyna/go/scyna/testing"
)

func TestCreateShouldReturnSuccess_Old(t *testing.T) {
	cleanup()
	requestCreate := &proto.CreateUserRequest{
		User: &proto.User{
			Email:    "a@gmail.com",
			Name:     "Nguyen Van A",
			Password: "1234565",
		},
	}

	scyna_test.CallServiceCheckCode(t, user.CREATE_USER_URL, requestCreate, 200)
	scyna_test.CallServiceCheckCode(t, user.GET_USER_URL, &proto.GetUserRequest{Email: "a@gmail.com"}, 200)
}

func TestCreateBadEmail_Old(t *testing.T) {
	cleanup()
	requestCreate := &proto.CreateUserRequest{
		User: &proto.User{
			Email:    "a+gmail.com",
			Name:     "Nguyen Van A",
			Password: "1234565",
		},
	}

	scyna_test.CallServiceCheckCode(t, user.CREATE_USER_URL, requestCreate, 400)
}