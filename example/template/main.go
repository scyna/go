package main

import (
	"flag"

	"template/event"
	"template/model"
	"template/service"

	"github.com/scyna/go/scyna"
)

const MODULE_CODE = "vf_support"

func main() {
	secret_ := flag.String("password", "123456789aA@#", "Authenticate By Token")
	managerUrl := flag.String("managerUrl", "http://127.0.0.1:8081", "Manager Url")
	flag.Parse()

	scyna.RemoteInit(scyna.RemoteConfig{
		ManagerUrl: *managerUrl,
		Name:       MODULE_CODE,
		Secret:     *secret_,
	})
	defer scyna.Release()

	scyna.UseRemoteLog(3)

	scyna.RegisterService(model.CREATE_NAME_URL, service.CreateName)

	scyna.RegisterEvent(MODULE_CODE, model.NAME_CREATED_CHANNEL, event.NameCreatedHandler)

	scyna.Start()
}
