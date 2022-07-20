package main

import (
	"github.com/scyna/go/manager/utils"
	"github.com/scyna/go/scyna"
)

const MODULE_CODE = "vf_op_init"

func main() {
	scyna.RemoteInit(scyna.RemoteConfig{
		ManagerUrl: "https://localhost:8081",
		Name:       MODULE_CODE,
		Secret:     "123456",
	})
	defer scyna.Release()

	utils.DeleteStream("vf_account")
	utils.CreateStreamForModule("vf_account")
	utils.CreateSyncConsumer2("vf_account", "account", "loyalty")
	utils.CreateSyncConsumer2("vf_account", "account", "salesforce")
	scyna.Start()
}
