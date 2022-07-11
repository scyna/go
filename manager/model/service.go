package model

import proto "github.com/scyna/go/manager/.proto/generated"

type Service struct {
	Module      string `db:"module_code"`
	URL         string `db:"url"`
	Description string `db:"description"`
}

func (s *Service) FromDTO(o *proto.Service) {
	s.URL = o.Url
	s.Module = o.Module
	s.Description = o.Description
}

func (s *Service) ToDTO() *proto.Service {
	return &proto.Service{
		Module:      s.Module,
		Url:         s.URL,
		Description: s.Description,
	}
}
