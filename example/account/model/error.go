package model

import "github.com/scyna/go/scyna"

var (
	ACCOUNT_EXISTS    = &scyna.Error{Code: 101, Message: "Account Exists"}
	ACCOUNT_NOT_FOUND = &scyna.Error{Code: 101, Message: "Account Not Found"}
	BAD_EMAIL         = &scyna.Error{Code: 102, Message: "Bad Email"}
	BAD_NAME          = &scyna.Error{Code: 103, Message: "Bad Name"}
	BAD_PASSWORD      = &scyna.Error{Code: 104, Message: "Bad Password"}
)
