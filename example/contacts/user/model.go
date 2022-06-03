package user

import "github.com/scyna/go/example/contacts/proto"

type User struct {
	ID       uint64 `db:"id"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	Password string `db:"password"`
}

func FromDTO(user *proto.User) *User {
	return &User{
		ID:       user.Id,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}

func (u *User) ToDTO() *proto.User {
	return &proto.User{
		Id:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}
}
