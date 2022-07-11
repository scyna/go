package model

import proto "github.com/scyna/go/manager/.proto/generated"

type Application struct {
	Organization     string `db:"org_code"`
	Code             string `db:"code"`
	Name             string `db:"name"`
	AuthenticatePath string `db:"auth"`
}

func (a *Application) FromDTO(o *proto.Application) {
	a.Code = o.Code
	a.Name = o.Name
	a.Organization = o.OrgCode
	a.AuthenticatePath = o.AuthPath
}

func (a *Application) ToDTO() *proto.Application {
	return &proto.Application{
		Code:     a.Code,
		Name:     a.Name,
		OrgCode:  a.Organization,
		AuthPath: a.AuthenticatePath,
	}
}
