package user

import (
	"github.com/scyna/go/example/scylla/proto"
	"github.com/scyna/go/scyna"
)

var (
	USER_EXISTED     = &scyna.Error{Code: 100, Message: "User Existed"}
	USER_NOT_EXISTED = &scyna.Error{Code: 101, Message: "User NOT Existed"}
)

type User struct {
	ID       uint64
	Name     string
	Email    string
	Password string
}

type IRepository interface {
	Create(LOG scyna.Logger, user *User) *scyna.Error
	Exist(LOG scyna.Logger, emai string) *scyna.Error
	GetByEmail(LOG scyna.Logger, email string) (*scyna.Error, *User)
	Release()
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
