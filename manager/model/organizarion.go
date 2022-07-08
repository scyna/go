package model

import proto "github.com/scyna/go/manager/.proto/generated"

type Organization struct {
	Code     string
	Name     string
	Password string
}

func (o *Organization) FromDTO(oganizarion *proto.Organization) {
	/*TODO*/
}

func (o *Organization) ToDTO() *proto.Organization {
	/*TODO*/
	return nil
}
