package scyna_test

import (
	"log"

	"github.com/scyna/go/scyna"
)

func Init() {
	scyna.RemoteInit(scyna.RemoteConfig{
		ManagerUrl: "http://127.0.0.1:8081",
		Name:       "scyna.test",
		Secret:     "123456",
	})
	log.Print(scyna.Session.ID())
	scyna.UseDirectLog(1)
}

func Release() {
	scyna.Release()
}
