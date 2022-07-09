package model

import proto "github.com/scyna/go/manager/.proto/generated"

type Service struct {
	Module      string
	URL         string
	Description string
}

func (s *Service) FromDTO(o *proto.Service) {
	/*TODO*/
}

func (m *Service) ToDTO() *proto.Service {
	/*TODO*/
	return nil
}
