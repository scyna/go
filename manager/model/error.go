package model

import "github.com/scyna/go/scyna"

var (
	ORGANIZATION_NOT_EXIST  = &scyna.Error{Code: 20, Message: "Organization Not Exist"}
	MODULE_EXISTED          = &scyna.Error{Code: 21, Message: "Module Existed"}
	MODULE_NOT_EXIST        = &scyna.Error{Code: 22, Message: "Module Not Exist"}
	CAN_NOT_CREATE_STREAM   = &scyna.Error{Code: 23, Message: "Can Not Create Stream"}
	CAN_NOT_CREATE_CONSUMER = &scyna.Error{Code: 24, Message: "Can Not Create Consumer"}
)
