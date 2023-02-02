package service

import (
	"fmt"
	proto "github.com/scyna/go/example/account/.proto/generated"
	"github.com/scyna/go/scyna"
)

func CreateAccount(s *scyna.Service, request *proto.CreateAccountRequest) {
	s.Info(fmt.Sprintf("Receive Create account request: %s", request))

}
