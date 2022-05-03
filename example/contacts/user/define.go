package user

import (
	"github.com/scyna/go/example/contacts/proto"
	"github.com/scyna/go/scyna"
)

var (
	USER_EXISTED     = &scyna.Error{Code: 100, Message: "User Existed"}
	USER_NOT_EXISTED = &scyna.Error{Code: 101, Message: "User NOT Existed"}
)

const (
	CREATE_USER_URL = "/example/user/create"
	GET_USER_URL    = "/example/user/get"
)

type IRepository interface {
	Create(LOG scyna.Logger, user *User) *scyna.Error
	Exist(LOG scyna.Logger, email string) *scyna.Error
	GetByEmail(LOG scyna.Logger, email string) (*scyna.Error, *User)
}

type User struct {
	ID       uint64 `db:"id"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	Password string `db:"password"`
}

func FromProto(user *proto.User) *User {
	return &User{
		ID:       user.Id,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}

func (u *User) ToProto() *proto.User {
	return &proto.User{
		Id:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}
}
