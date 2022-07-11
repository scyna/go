package model

import "github.com/scyna/go/scyna"

var (
	ORGANIZATION_NOT_EXIST  = &scyna.Error{Code: 20, Message: "Organization Not Exist"}
	ORGANIZATION_EXISTED    = &scyna.Error{Code: 21, Message: "Organization Existed"}
	APPLICATION_NOT_EXIST   = &scyna.Error{Code: 20, Message: "Application Not Exist"}
	APPLICATION_EXISTED     = &scyna.Error{Code: 21, Message: "Application Existed"}
	MODULE_EXISTED          = &scyna.Error{Code: 22, Message: "Module Existed"}
	SERVICE_EXISTED         = &scyna.Error{Code: 22, Message: "Service Existed"}
	CLIENT_EXISTED          = &scyna.Error{Code: 22, Message: "Client Existed"}
	MODULE_NOT_EXIST        = &scyna.Error{Code: 23, Message: "Module Not Exist"}
	CAN_NOT_CREATE_STREAM   = &scyna.Error{Code: 24, Message: "Can Not Create Stream"}
	CAN_NOT_CREATE_CONSUMER = &scyna.Error{Code: 25, Message: "Can Not Create Consumer"}
	SERVICE_PATH_BAD_FORMAT = &scyna.Error{Code: 25, Message: "Service Path Bad Format"}
)
