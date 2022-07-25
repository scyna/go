package main

import (
	"flag"

	"github.com/scyna/go/manager/model"
	"github.com/scyna/go/manager/service/application"
	"github.com/scyna/go/manager/service/client"
	"github.com/scyna/go/manager/service/event"
	"github.com/scyna/go/manager/service/module"
	"github.com/scyna/go/manager/service/organization"
	"github.com/scyna/go/manager/service/service"
	"github.com/scyna/go/manager/service/sync"
	"github.com/scyna/go/manager/service/task"
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

	scyna.RegisterService(model.ORGANIZATION_CREATE_URL, organization.CreateOrganization)
	scyna.RegisterService(model.ORGANIZATION_DESTROY_URL, organization.DestroyOrganization)
	scyna.RegisterService(model.APPLICATION_CREATE_URL, application.CreateApplication)
	scyna.RegisterService(model.MODULE_CREATE_URL, module.CreateModule)
	scyna.RegisterService(model.SERVICE_CREATE_URL, service.CreateService)
	scyna.RegisterService(model.CLIENT_CREATE_URL, client.CreateClient)
	scyna.RegisterService(model.CLIENT_ADD_SERVICE_URL, client.AddService)
	scyna.RegisterService(model.CLIENT_REMOVE_SERVICE_URL, client.RemoveService)
	scyna.RegisterService(model.EVENT_CREATE_URL, event.CreateEvent)
	scyna.RegisterService(model.SYNC_CREATE_URL, sync.CreateSync)
	scyna.RegisterService(model.TASK_CREATE_URL, task.CreateTask)
	scyna.Start()
}
