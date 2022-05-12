package user

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/scyna/go/example/contacts/proto"
	"github.com/scyna/go/scyna"
)

func Create(ctx *scyna.Service, request *proto.CreateUserRequest) {
	ctx.Logger.Info("Receive CreateUserRequest")
	if err := validateCreateRequest(request.User); err != nil {
		ctx.Error(scyna.REQUEST_INVALID)
		return
	}

	if err, _ := Repository.GetByEmail(ctx.Logger, request.User.Email); err == nil {
		ctx.Error(USER_EXISTED)
		return
	}

	user := FromDTO(request.User)
	user.ID = scyna.ID.Next()
	if err := Repository.Create(ctx.Logger, user); err != nil {
		ctx.Error(err)
		return
	}
	ctx.Done(&proto.CreateUserResponse{Id: user.ID})
}

func validateCreateRequest(user *proto.User) error {
	return validation.ValidateStruct(user,
		validation.Field(&user.Email, validation.Required, is.Email),
		validation.Field(&user.Password, validation.Required, validation.Length(5, 10)),
		validation.Field(&user.Name, validation.Required, validation.Length(1, 100)))
}
