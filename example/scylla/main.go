package main

import (
	"github.com/scyna/go/example/scylla/user"
	"github.com/scyna/go/scyna"
)

const MODULE_CODE = "scyna.example"

func main() {
	scyna.RemoteInit(scyna.RemoteConfig{
		ManagerUrl: "http://localhost:8081",
		Name:       MODULE_CODE,
		Secret:     "123456",
	})
	defer scyna.Release()

	scyna.RegisterService("/scyna.example/user/create", user.Create)

	scyna.Start()
}
