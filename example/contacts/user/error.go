package user

import (
	"github.com/scyna/go/scyna"
)

var (
	USER_EXISTED     = &scyna.Error{Code: 100, Message: "User Existed"}
	USER_NOT_EXISTED = &scyna.Error{Code: 101, Message: "User NOT Existed"}
)
