package model

import (
	proto "github.com/scyna/go/manager/.proto/generated"
)

type Module struct {
	Organization string `db:"org_code"`
	Code         string `db:"code"`
	Description  string `db:"description"`
	Secret       string `db:"secret"`
}

func (m *Module) FromDTO(o *proto.Module) {
	m.Code = o.Code
	m.Organization = o.Organization
	m.Description = o.Description
	m.Secret = o.Secret
}

func (m *Module) ToDTO() *proto.Module {
	return &proto.Module{
		Organization: m.Organization,
		Code:         m.Code,
		Description:  m.Description,
		Secret:       "", // NO RETURN
	}
}
