package user

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/scyna/go/example/contacts/proto"
	"github.com/scyna/go/scyna"
)

func CreateUser(s *scyna.Service, request *proto.User) {
	s.Logger.Info("Receive CreateUserRequest")
	if err := validateCreateUserRequest(request); err != nil {
		s.Error(scyna.REQUEST_INVALID)
		return
	}

	if err, _ := Repository.GetByEmail(s.Logger, request.Email); err == nil {
		s.Error(USER_EXISTED)
		return
	}

	user := FromDTO(request)
	user.ID = scyna.ID.Next()
	//if err := Repository.Create(s.Logger, user); err != nil {
	//	s.Error(err)
	//	return
	//}

	//s.PostSync("account", user.ToDTO())

	s.Done(&proto.CreateUserResponse{Id: user.ID})
}

func validateCreateUserRequest(user *proto.User) error {
	return validation.ValidateStruct(user,
		validation.Field(&user.Email, validation.Required, is.Email),
		validation.Field(&user.Password, validation.Required, validation.Length(5, 10)),
		validation.Field(&user.Name, validation.Required, validation.Length(1, 100)))
}
