package model

import proto "github.com/scyna/go/manager/.proto/generated"

type Application struct {
	Organization     string `db:"org_code"`
	Code             string `db:"code"`
	Name             string `db:"name"`
	AuthenticatePath string `db:"auth"`
}

func (c *Application) FromDTO(o *proto.Application) {
	/*TODO*/
}

func (c *Application) ToDTO() *proto.Application {
	/*TODO*/
	return nil
}
