package service

import (
	"fmt"
	proto "github.com/scyna/go/example/account/.proto/generated"
	"github.com/scyna/go/example/account/repository"
	"github.com/scyna/go/scyna"
)

func GetAccountByEmail(s *scyna.Service, request *proto.GetAccountByEmailRequest) {
	s.Info(fmt.Sprintf("Receive Get account by email request: %s", request))

	account, err := repository.GetAccountByEmail(s.Logger, request.Email)
	if err != nil {
		s.Logger.Error("Create account err: " + err.Message)
		s.Error(err)
		return
	}

	s.Done(&proto.GetAccountResponse{
		Id:    account.ID,
		Email: account.Email.Value,
		Name:  account.Name.Value,
	})

}
