package main

import (
	"github.com/scyna/go/example/contacts/user"
	"github.com/scyna/go/scyna"
)

const MODULE_CODE = "nats"

func main() {
	scyna.RemoteInit(scyna.RemoteConfig{
		ManagerUrl: "https://localhost:8081",
		Name:       MODULE_CODE,
		Secret:     "123456789aA@#",
	})
	defer scyna.Release()

	scyna.RegisterService("/scyna.example/user/create", user.CreateUser)
	scyna.RegisterSync("account", "receiver1", user.HandlerSyncMessage)
	scyna.Start()
}
