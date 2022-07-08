package model

import proto "github.com/scyna/go/manager/.proto/generated"

type Module struct {
	Organization string
	Code         string
	Description  string
	Secret       string
}

func (m *Module) FromDTO(o *proto.Module) {
	/*TODO*/
}

func (m *Module) ToDTO() *proto.Module {
	/*TODO*/
	return nil
}
