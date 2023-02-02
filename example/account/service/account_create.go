package service

import (
	"fmt"
	proto "github.com/scyna/go/example/account/.proto/generated"
	"github.com/scyna/go/example/account/model"
	"github.com/scyna/go/example/account/repository"
	"github.com/scyna/go/scyna"
)

func CreateAccount(s *scyna.Service, request *proto.CreateAccountRequest) {
	s.Info(fmt.Sprintf("Receive Create account request: %s", request))

	var account = model.Account{
		ID: scyna.ID.Next(),
		Name: model.Name{
			Value: request.Name,
		},
		Email: model.EmailAddress{
			Value: request.Email,
		},
		Password: model.Password{
			Value: request.Password,
		},
	}

	if err := repository.CreateAccount(s.Logger, account); err != nil {
		s.Logger.Error("Create account err: " + err.Message)
		s.Error(err)
		return
	}

	s.Done(&proto.CreateAccountResponse{
		Id: account.ID,
	})

	s.PostEventAndActivity(model.ACCOUNT_CREATED_CHANNEL, &proto.AccountCreated{
		Id:    account.ID,
		Email: account.Email.Value,
		Name:  request.Name,
	}, []uint64{
		account.ID,
	})

}
