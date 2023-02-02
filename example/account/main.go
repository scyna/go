package main

import (
	"flag"
	"github.com/scyna/go/example/account/event"
	"github.com/scyna/go/example/account/model"
	"github.com/scyna/go/example/account/service"

	"github.com/scyna/go/scyna"
)

const MODULE_CODE = "account"

func main() {
	secret_ := flag.String("password", "123456789aA@#", "Authenticate By Token")
	managerUrl := flag.String("managerUrl", "http://127.0.0.1:8081", "Manager Url")
	flag.Parse()

	scyna.RemoteInit(scyna.RemoteConfig{
		ManagerUrl: *managerUrl,
		Name:       MODULE_CODE,
		Secret:     *secret_,
	})
	scyna.UseRemoteLog(3)

	scyna.RegisterService(model.CREATE_ACCOUNT_URL, service.CreateAccount)
	scyna.RegisterService(model.GET_ACCOUNT_BY_EMAIl_URL, service.GetAccountByEmail)

	scyna.RegisterEvent(MODULE_CODE, model.ACCOUNT_CREATED_CHANNEL, event.AccountCreatedHandler)

	defer scyna.Release()

	scyna.Start()
}
