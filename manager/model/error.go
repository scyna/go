package model

import "github.com/scyna/go/scyna"

var (
	ORGANIZATION_NOT_EXIST = &scyna.Error{Code: 20, Message: "Organization not exist"}
	MODULE_EXISTED         = &scyna.Error{Code: 21, Message: "Module existed"}
)
