package main

import (
	"flag"
	"github.com/scyna/go/manager/model"
	"github.com/scyna/go/manager/service/application"
	"github.com/scyna/go/manager/service/organization"
	"github.com/scyna/go/scyna"
)

const MODULE_CODE = "scyna.manager"

func main() {
	secret_ := flag.String("password", "123456", "AuthenticateByToken")
	managerUrl := flag.String("managerUrl", "https://127.0.0.1:8081", "Manager Url")
	flag.Parse()

	scyna.RemoteInit(scyna.RemoteConfig{
		ManagerUrl: *managerUrl,
		Name:       MODULE_CODE,
		Secret:     *secret_,
	})
	scyna.UseDirectLog(3)
	defer scyna.Release()

	scyna.RegisterService(model.MANAGER_CREATE_ORGANIZATION_URL, organization.CreateOrganization)
	scyna.RegisterService(model.MANAGER_DESTROY_ORGANIZATION_URL, organization.DestroyOrganization)
	scyna.RegisterService(model.MANAGER_CREATE_APPLICATION_URL, application.CreateApplication)
	scyna.Start()
}
