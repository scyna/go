package model

import proto "github.com/scyna/go/manager/.proto/generated"

type Organization struct {
	Code     string
	Name     string
	Password string
}

func (org *Organization) FromDTO(o *proto.Organization) {
	/*TODO*/
}

func (org *Organization) ToDTO() *proto.Organization {
	/*TODO*/
	return nil
}
