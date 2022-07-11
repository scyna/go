package model

import proto "github.com/scyna/go/manager/.proto/generated"

type Organization struct {
	Code     string
	Name     string
	Password string
}

func (o *Organization) FromDTO(organization *proto.Organization) {
	o.Code = organization.Code
	o.Name = organization.Name
	o.Password = organization.Password
}

func (o *Organization) ToDTO() *proto.Organization {
	return &proto.Organization{
		Code:     o.Code,
		Name:     o.Name,
		Password: "", // NO RETURN
	}
}
