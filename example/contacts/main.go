package main

import (
	"github.com/scyna/go/example/contacts/user"
	"github.com/scyna/go/scyna"
)

const MODULE_CODE = "nats"

func main() {
	scyna.RemoteInit(scyna.RemoteConfig{
		Url:     "https://localhost:8081",
		Context: MODULE_CODE,
		Secret:  "123456789aA@#",
	})
	defer scyna.Release()

	scyna.RegisterService("/scyna.example/user/create", user.CreateUser)
	scyna.RegisterSync("channel", "receiver", user.HandlerSyncMessage)
	scyna.Start()
}
