package main

import (
	"github.com/scyna/go/example/contacts/user"
	"github.com/scyna/go/scyna"
)

//const MODULE_CODE = "vf_profile"

const MODULE_CODE = "vf_account"

func main() {
	scyna.RemoteInit(scyna.RemoteConfig{
		ManagerUrl: "https://localhost:8081",
		Name:       MODULE_CODE,
		Secret:     "123456789aA@#",
	})
	defer scyna.Release()

	scyna.RegisterService("/scyna.example/user/create", user.CreateUser)
	//scyna.RegisterEvent("vf_account", "account_loyalty", user.HandlerEventMessage)
	scyna.RegisterSync("account", "salesforce", user.HandlerSyncMessage)
	scyna.Start()
}
