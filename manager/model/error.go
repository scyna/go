package model

import "github.com/scyna/go/scyna"

var (
	ORGANIZATION_NOT_EXIST = &scyna.Error{Code: 20, Message: "Organization Not Exist"}
	MODULE_EXISTED         = &scyna.Error{Code: 21, Message: "Module Existed"}
	CAN_NOT_CREATE_STREAM  = &scyna.Error{Code: 22, Message: "Can Not Create Stream"}
)
